package store

import (
	"encoding/binary"
	"fmt"
	"path/filepath"

	"github.com/coreos/bbolt"
	"github.com/gaia-pipeline/gaia"
)

var (
	// Name of the bucket where we store user objects
	userBucket = []byte("Users")

	// Name of the bucket where we store information about pipelines
	pipelineBucket = []byte("Pipelines")

	// Name of the bucket where we store information about pipelines
	// which are not yet compiled (create pipeline)
	createPipelineBucket = []byte("CreatePipelines")

	// Name of the bucket where we store all pipeline runs.
	pipelineRunBucket = []byte("PipelineRun")

	// Where we store all the created permission groups
	permGroupsBucket = []byte("PermissionGroups")

	// Where we store all users permissions
	userPermsBucket = []byte("UsersPermissions")
)

const (
	// Username and password of the first admin user
	adminUsername = "admin"
	adminPassword = "admin"

	// Bolt database file name
	boltDBFileName = "gaia.db"
)

// BoltStore represents the access type for store
type BoltStore struct {
	db *bolt.DB
}

// GaiaStore is the interface that defines methods needed to store
// pipeline and user related information.
type GaiaStore interface {
	Init() error
	CreatePipelinePut(createPipeline *gaia.CreatePipeline) error
	CreatePipelineGet() (listOfPipelines []gaia.CreatePipeline, err error)
	PipelinePut(pipeline *gaia.Pipeline) error
	PipelineGet(id int) (pipeline *gaia.Pipeline, err error)
	PipelineGetByName(name string) (pipline *gaia.Pipeline, err error)
	PipelineGetRunHighestID(pipeline *gaia.Pipeline) (id int, err error)
	PipelinePutRun(r *gaia.PipelineRun) error
	PipelineGetScheduled(limit int) ([]*gaia.PipelineRun, error)
	PipelineGetRunByPipelineIDAndID(pipelineid int, runid int) (*gaia.PipelineRun, error)
	PipelineGetAllRuns(pipelineID int) ([]gaia.PipelineRun, error)
	PipelineGetLatestRun(pipelineID int) (*gaia.PipelineRun, error)
	PipelineDelete(id int) error
	UserPut(u *gaia.User, encryptPassword bool) error
	UserAuth(u *gaia.User, updateLastLogin bool) (*gaia.User, error)
	UserGet(username string) (*gaia.User, error)
	UserGetAll() ([]gaia.User, error)
	UserDelete(u string) error
	UserPermissionsPut(perms *gaia.UserPermissions) error
	UserPermissionsGet(username string) (*gaia.UserPermissions, error)
	PermissionGroupGetAll() ([]*gaia.PermissionGroup, error)
	PermissionGroupGet(name string) (*gaia.PermissionGroup, error)
	PermissionGroupPut(group *gaia.PermissionGroup) error
	PermissionGroupCreate(group *gaia.PermissionGroup) error
	PermissionGroupDelete(name string) error
}

// Compile time interface compliance check for BoltStore. If BoltStore
// wouldn't implement GaiaStore this line wouldn't compile.
var _ GaiaStore = (*BoltStore)(nil)

// NewBoltStore creates a new instance of Store.
func NewBoltStore() *BoltStore {
	s := &BoltStore{}

	return s
}

// Init creates the data folder if not exists,
// generates private key and bolt database.
// This should be called only once per database
// because bolt holds a lock on the database file.
func (s *BoltStore) Init() error {
	// Open connection to bolt database
	path := filepath.Join(gaia.Cfg.DataPath, boltDBFileName)
	db, err := bolt.Open(path, gaia.Cfg.Bolt.Mode, nil)
	if err != nil {
		return err
	}
	s.db = db

	// Setup database
	return s.setupDatabase()
}

// setupDatabase create all buckets in the db.
// Additionally, it makes sure that the admin user exists.
func (s *BoltStore) setupDatabase() error {
	// Create bucket if not exists function
	var bucketName []byte
	c := func(tx *bolt.Tx) error {
		_, err := tx.CreateBucketIfNotExists(bucketName)
		if err != nil {
			return fmt.Errorf("create bucket: %s", err)
		}
		return nil
	}

	// Make sure buckets exist
	bucketName = userBucket
	err := s.db.Update(c)
	if err != nil {
		return err
	}
	bucketName = userPermsBucket
	err = s.db.Update(c)
	if err != nil {
		return err
	}
	bucketName = permGroupsBucket
	err = s.db.Update(c)
	if err != nil {
		return err
	}
	bucketName = pipelineBucket
	err = s.db.Update(c)
	if err != nil {
		return err
	}
	bucketName = createPipelineBucket
	err = s.db.Update(c)
	if err != nil {
		return err
	}
	bucketName = pipelineRunBucket
	err = s.db.Update(c)
	if err != nil {
		return err
	}

	// Make sure that the user "admin" does exist
	admin, err := s.UserGet(adminUsername)
	if err != nil {
		return err
	}

	// Create admin user if we cannot find it
	if admin == nil {
		err = s.UserPut(&gaia.User{
			DisplayName: adminUsername,
			Username:    adminUsername,
			Password:    adminPassword,
		}, true)

		if err != nil {
			return err
		}
	}

	return nil
}

// itob returns an 8-byte big endian representation of v.
func itob(v int) []byte {
	b := make([]byte, 8)
	binary.BigEndian.PutUint64(b, uint64(v))
	return b
}
