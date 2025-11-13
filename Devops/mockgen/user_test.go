package mockgen

//
//import (
//	"testing"
//
//	"github.com/golang/mock/gomock"
//	"github.com/stretchr/testify/assert"
//)
//
//func TestSaveUser(t *testing.T) {
//	ctrl := gomock.NewController(t)
//	defer ctrl.Finish()
//
//	mockDB := NewMockDatabase(ctrl)
//
//	mockDB.EXPECT().Save("user1", "John").Return(nil)
//
//	err := SaveUser(mockDB, "user1", "John")
//
//	assert.NoError(t, err)
//}
//
//func TestGetUser(t *testing.T) {
//	ctrl := gomock.NewController(t)
//	defer ctrl.Finish()
//
//	mockDB := NewMockDatabase(ctrl)
//
//	// Define what we expect the mock to return
//	mockDB.EXPECT().Load("user1").Return("John", nil)
//
//	name, err := GetUser(mockDB, "user1")
//
//	assert.NoError(t, err)
//	assert.Equal(t, "John", name)
//}
