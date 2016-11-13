package kiasu

// AUTOGENERATED BY MOQ
// github.com/matryer/moq

// PrimitiveStoreMock is a mock implementation of PrimitiveStore.
//
//     func TestSomethingThatUsesPrimitiveStore(t *testing.T) {
//
//         // make and configure a mocked PrimitiveStore
//         mockedPrimitiveStore := &PrimitiveStoreMock{
//             GetFeedFunc: func(id string) (*Feed, error) {
// 	               panic("TODO: mock out the GetFeed function")
//             },
//             GetFeedsFunc: func(pg *Pagination) ([]Feed, error) {
// 	               panic("TODO: mock out the GetFeeds function")
//             },
//             GetPostFunc: func(id string) (*Post, error) {
// 	               panic("TODO: mock out the GetPost function")
//             },
//             GetPostsFunc: func(feedID string, pg *Pagination) ([]Post, error) {
// 	               panic("TODO: mock out the GetPosts function")
//             },
//             GetReadStatusFunc: func(id string) (*ReadStatus, error) {
// 	               panic("TODO: mock out the GetReadStatus function")
//             },
//             GetReadStatusByPostIDFunc: func(postID string, userID string) (*ReadStatus, error) {
// 	               panic("TODO: mock out the GetReadStatusByPostID function")
//             },
//             GetSessionFunc: func(id string) (*Session, error) {
// 	               panic("TODO: mock out the GetSession function")
//             },
//             GetSessionByAccessTokenFunc: func(token string) (*Session, error) {
// 	               panic("TODO: mock out the GetSessionByAccessToken function")
//             },
//             GetSessionsByUserIDFunc: func(userID string, pg *Pagination) ([]Session, error) {
// 	               panic("TODO: mock out the GetSessionsByUserID function")
//             },
//             GetUserFunc: func(id string) (*User, error) {
// 	               panic("TODO: mock out the GetUser function")
//             },
//             SaveFeedFunc: func(in1 *Feed) (*Feed, error) {
// 	               panic("TODO: mock out the SaveFeed function")
//             },
//             SavePostFunc: func(in1 *Post) (*Post, error) {
// 	               panic("TODO: mock out the SavePost function")
//             },
//             SaveReadStatusFunc: func(in1 *ReadStatus) (*ReadStatus, error) {
// 	               panic("TODO: mock out the SaveReadStatus function")
//             },
//             SaveSessionFunc: func(in1 *Session) (*User, error) {
// 	               panic("TODO: mock out the SaveSession function")
//             },
//             SaveUserFunc: func(in1 *User) (*User, error) {
// 	               panic("TODO: mock out the SaveUser function")
//             },
//         }
//s
//         // TODO: use mockedPrimitiveStore in code that requires PrimitiveStore
//
//     }
type PrimitiveStoreMock struct {
	// GetFeedFunc mocks the GetFeed function.
	GetFeedFunc func(id string) (*Feed, error)
	// GetFeedsFunc mocks the GetFeeds function.
	GetFeedsFunc func(pg *Pagination) ([]Feed, error)
	// GetPostFunc mocks the GetPost function.
	GetPostFunc func(fid, pid string) (*Post, error)
	// GetPostsFunc mocks the GetPosts function.
	GetPostsFunc func(feedID string, pg *Pagination) ([]Post, error)
	// GetReadStatusFunc mocks the GetReadStatus function.
	GetReadStatusFunc func(id string) (*ReadStatus, error)
	// GetReadStatusByPostIDFunc mocks the GetReadStatusByPostID function.
	GetReadStatusByPostIDFunc func(postID string, userID string) (*ReadStatus, error)
	// GetSessionFunc mocks the GetSession function.
	GetSessionFunc func(id string) (*Session, error)
	// GetSessionByAccessTokenFunc mocks the GetSessionByAccessToken function.
	GetSessionByAccessTokenFunc func(token string) (*Session, error)
	// GetSessionsByUserIDFunc mocks the GetSessionsByUserID function.
	GetSessionsByUserIDFunc func(userID string, pg *Pagination) ([]Session, error)
	// GetUserFunc mocks the GetUser function.
	GetUserFunc func(id string) (*User, error)
	// SaveFeedFunc mocks the SaveFeed function.
	SaveFeedFunc func(in1 *Feed) (*Feed, error)
	// SavePostFunc mocks the SavePost function.
	SavePostFunc func(in1 *Post) (*Post, error)
	// SaveReadStatusFunc mocks the SaveReadStatus function.
	SaveReadStatusFunc func(in1 *ReadStatus) (*ReadStatus, error)
	// SaveSessionFunc mocks the SaveSession function.
	SaveSessionFunc func(in1 *Session) (*User, error)
	// SaveUserFunc mocks the SaveUser function.
	SaveUserFunc func(in1 *User) (*User, error)
}

// GetFeed calls GetFeedFunc.
func (mock *PrimitiveStoreMock) GetFeed(id string) (*Feed, error) {
	if mock.GetFeedFunc == nil {
		panic("moq: GetFeedFunc is nil but was just called")
	}
	return mock.GetFeedFunc(id)
}

// GetFeeds calls GetFeedsFunc.
func (mock *PrimitiveStoreMock) GetFeeds(pg *Pagination) ([]Feed, error) {
	if mock.GetFeedsFunc == nil {
		panic("moq: GetFeedsFunc is nil but was just called")
	}
	return mock.GetFeedsFunc(pg)
}

// GetPost calls GetPostFunc.
func (mock *PrimitiveStoreMock) GetPost(feedID, postID string) (*Post, error) {
	if mock.GetPostFunc == nil {
		panic("moq: GetPostFunc is nil but was just called")
	}
	return mock.GetPostFunc(feedID, postID)
}

// GetPosts calls GetPostsFunc.
func (mock *PrimitiveStoreMock) GetPosts(feedID string, pg *Pagination) ([]Post, error) {
	if mock.GetPostsFunc == nil {
		panic("moq: GetPostsFunc is nil but was just called")
	}
	return mock.GetPostsFunc(feedID, pg)
}

// GetReadStatus calls GetReadStatusFunc.
func (mock *PrimitiveStoreMock) GetReadStatus(id string) (*ReadStatus, error) {
	if mock.GetReadStatusFunc == nil {
		panic("moq: GetReadStatusFunc is nil but was just called")
	}
	return mock.GetReadStatusFunc(id)
}

// GetReadStatusByPostID calls GetReadStatusByPostIDFunc.
func (mock *PrimitiveStoreMock) GetReadStatusByPostID(postID string, userID string) (*ReadStatus, error) {
	if mock.GetReadStatusByPostIDFunc == nil {
		panic("moq: GetReadStatusByPostIDFunc is nil but was just called")
	}
	return mock.GetReadStatusByPostIDFunc(postID, userID)
}

// GetSession calls GetSessionFunc.
func (mock *PrimitiveStoreMock) GetSession(id string) (*Session, error) {
	if mock.GetSessionFunc == nil {
		panic("moq: GetSessionFunc is nil but was just called")
	}
	return mock.GetSessionFunc(id)
}

// GetSessionByAccessToken calls GetSessionByAccessTokenFunc.
func (mock *PrimitiveStoreMock) GetSessionByAccessToken(token string) (*Session, error) {
	if mock.GetSessionByAccessTokenFunc == nil {
		panic("moq: GetSessionByAccessTokenFunc is nil but was just called")
	}
	return mock.GetSessionByAccessTokenFunc(token)
}

// GetSessionsByUserID calls GetSessionsByUserIDFunc.
func (mock *PrimitiveStoreMock) GetSessionsByUserID(userID string, pg *Pagination) ([]Session, error) {
	if mock.GetSessionsByUserIDFunc == nil {
		panic("moq: GetSessionsByUserIDFunc is nil but was just called")
	}
	return mock.GetSessionsByUserIDFunc(userID, pg)
}

// GetUser calls GetUserFunc.
func (mock *PrimitiveStoreMock) GetUser(id string) (*User, error) {
	if mock.GetUserFunc == nil {
		panic("moq: GetUserFunc is nil but was just called")
	}
	return mock.GetUserFunc(id)
}

// SaveFeed calls SaveFeedFunc.
func (mock *PrimitiveStoreMock) SaveFeed(in1 *Feed) (*Feed, error) {
	if mock.SaveFeedFunc == nil {
		panic("moq: SaveFeedFunc is nil but was just called")
	}
	return mock.SaveFeedFunc(in1)
}

// SavePost calls SavePostFunc.
func (mock *PrimitiveStoreMock) SavePost(in1 *Post) (*Post, error) {
	if mock.SavePostFunc == nil {
		panic("moq: SavePostFunc is nil but was just called")
	}
	return mock.SavePostFunc(in1)
}

// SaveReadStatus calls SaveReadStatusFunc.
func (mock *PrimitiveStoreMock) SaveReadStatus(in1 *ReadStatus) (*ReadStatus, error) {
	if mock.SaveReadStatusFunc == nil {
		panic("moq: SaveReadStatusFunc is nil but was just called")
	}
	return mock.SaveReadStatusFunc(in1)
}

// SaveSession calls SaveSessionFunc.
func (mock *PrimitiveStoreMock) SaveSession(in1 *Session) (*User, error) {
	if mock.SaveSessionFunc == nil {
		panic("moq: SaveSessionFunc is nil but was just called")
	}
	return mock.SaveSessionFunc(in1)
}

// SaveUser calls SaveUserFunc.
func (mock *PrimitiveStoreMock) SaveUser(in1 *User) (*User, error) {
	if mock.SaveUserFunc == nil {
		panic("moq: SaveUserFunc is nil but was just called")
	}
	return mock.SaveUserFunc(in1)
}
