	user_model "code.gitea.io/gitea/models/user"
	user := unittest.AssertExistsAndLoadBean(t, &user_model.User{ID: 2}).(*user_model.User)
	user := unittest.AssertExistsAndLoadBean(t, &user_model.User{ID: 2}).(*user_model.User)
	user := unittest.AssertExistsAndLoadBean(t, &user_model.User{ID: 2}).(*user_model.User)
	user := unittest.AssertExistsAndLoadBean(t, &user_model.User{ID: 2}).(*user_model.User)
	user := unittest.AssertExistsAndLoadBean(t, &user_model.User{ID: 2}).(*user_model.User)
}

func TestGetFileHistory(t *testing.T) {
	defer prepareTestEnv(t)()
	user := unittest.AssertExistsAndLoadBean(t, &user_model.User{ID: 2}).(*user_model.User)
	// Login as User2.
	session := loginUser(t, user.Name)
	token := getTokenForLoggedInUser(t, session)

	req := NewRequestf(t, "GET", "/api/v1/repos/%s/repo16/commits?path=readme.md&token="+token+"&sha=good-sign", user.Name)
	resp := session.MakeRequest(t, req, http.StatusOK)

	var apiData []api.Commit
	DecodeJSON(t, resp, &apiData)
	assert.Len(t, apiData, 1)
	assert.Equal(t, "f27c2b2b03dcab38beaf89b0ab4ff61f6de63441", apiData[0].CommitMeta.SHA)
	compareCommitFiles(t, []string{"readme.md"}, apiData[0].Files)