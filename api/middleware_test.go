package api

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/Aashish-32/bank/token"
	"github.com/Aashish-32/bank/util"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/require"
)

func addAuthorization(t *testing.T,
	request *http.Request,
	tokenmaker token.Maker,
	authorizationType string,
	username string,
	duration time.Duration,

) {

	token, err := tokenmaker.CreateToken(username, duration)
	require.NoError(t, err)
	authorizationHeader := fmt.Sprintf("%s %s", authorizationType, token)
	request.Header.Set(authorizationHeaderKey, authorizationHeader)
}

func TestAuthMiddleware(t *testing.T) {

	//table driven tests

	testCases := []struct {
		name          string
		setupAuth     func(t *testing.T, req *http.Request, tokenmaker token.Maker)
		checkResponse func(t *testing.T, recorder *httptest.ResponseRecorder)
	}{

		{
			name: "ok",
			setupAuth: func(t *testing.T, req *http.Request, tokenmaker token.Maker) {

				addAuthorization(t, req, tokenmaker, authorizationTypeBearer, "user", time.Minute)
			},

			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusOK, recorder.Code)
			},
		},

		{
			name: "UnsupportedAuthorization",
			setupAuth: func(t *testing.T, req *http.Request, tokenmaker token.Maker) {

				addAuthorization(t, req, tokenmaker, util.RandomString(5), "user", time.Minute)
			},

			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusUnauthorized, recorder.Code)
			},
		},

		{
			name: "InvalidAuthorizationFormat",
			setupAuth: func(t *testing.T, req *http.Request, tokenmaker token.Maker) {

				addAuthorization(t, req, tokenmaker, "", "user", time.Minute)
			},

			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusUnauthorized, recorder.Code)
			},
		},

		{
			name: "eXPIREDTOKEN",
			setupAuth: func(t *testing.T, req *http.Request, tokenmaker token.Maker) {

				addAuthorization(t, req, tokenmaker, authorizationTypeBearer, "user", -time.Minute)
			},

			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusUnauthorized, recorder.Code)
			},
		},
	}

	for i := range testCases {
		tc := testCases[i]

		//generating subtest

		t.Run(tc.name, func(t *testing.T) {

			// os.Setenv("token_symmetric_key", util.RandomString(32))

			server := newTestServer(t, nil)

			authPath := "/auth"

			//setting up route

			server.router.GET(
				authPath,
				authMiddleware(server.tokenMaker),
				func(ctx *gin.Context) {
					ctx.JSON(http.StatusOK, gin.H{})

				},
			)
			recorder := httptest.NewRecorder()

			request, err := http.NewRequest(http.MethodGet, authPath, nil)
			require.NoError(t, err)

			tc.setupAuth(t, request, server.tokenMaker)
			server.router.ServeHTTP(recorder, request)
			tc.checkResponse(t, recorder)

		})

	}

}
