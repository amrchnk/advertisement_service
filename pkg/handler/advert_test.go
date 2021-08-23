package handler

import (
    "bytes"
    "testing"
    "github.com/gin-gonic/gin"
    "github.com/stretchr/testify/assert"
    "github.com/amrchnk/advertisement_service/pkg/models"
    "github.com/amrchnk/advertisement_service/pkg/service"
    mock_service "github.com/amrchnk/advertisement_service/pkg/service/mocks"
    "net/http/httptest"
    "github.com/golang/mock/gomock"
    "errors"
)

func TestHandler_createAdvert(t *testing.T){
    type mockBehavior func(s *mock_service.MockAdvert,advert models.Advert)

    testTable:=[]struct{
        name string
        inputBody string
        inputAdvert models.Advert
        mockBehavior mockBehavior
        expectedStatusCode int
        expectedRequestBody string
    }{
        {
            name:"OK",
            inputBody:`{"title":"Test title","description":"test description","price":111,"photos":["one","two","three"]}`,
            inputAdvert: models.Advert{
                Title:"Test title",
                Description:"test description",
                Price:111,
                Photos:[]string{"one","two","three"},
            },
            mockBehavior: func(s *mock_service.MockAdvert, advert models.Advert){
                s.EXPECT().CreateAdvert(advert).Return(1,nil)
            },
            expectedStatusCode:200,
            expectedRequestBody:`{"id":1,"status":200}`,
        },
        {
            name:"Empty fields",
            inputBody:`{"description":"test description","price":111,"photos":["one","two","three"]}`,
            mockBehavior: func(s *mock_service.MockAdvert, advert models.Advert){},
            expectedStatusCode:400,
            expectedRequestBody:`{"id":-1,"status":400}`,
        },
        {
            name:"Binding error",
            inputBody:`{"title":111,"description":"test description","price":111,"photos":["one","two","three"]}`,
            mockBehavior: func(s *mock_service.MockAdvert, advert models.Advert){},
            expectedStatusCode:400,
            expectedRequestBody:`{"id":-1,"status":400}`,
        },
        {
            name:"Response error",
            inputBody:`{"title":"Test title","description":"test description","price":111,"photos":["one","two","three"]}`,
            inputAdvert: models.Advert{
                Title:"Test title",
                Description:"test description",
                Price:111,
                Photos:[]string{"one","two","three"},
            },
            mockBehavior: func(s *mock_service.MockAdvert, advert models.Advert){
                s.EXPECT().CreateAdvert(advert).Return(0, errors.New("some server error"))
            },
            expectedStatusCode:500,
            expectedRequestBody:`{"id":-1,"status":500}`,
        },
    }

    for _,testCase:=range testTable{
        t.Run(testCase.name,func(t *testing.T){
            c:=gomock.NewController(t)
            defer c.Finish()

            adv:=mock_service.NewMockAdvert(c)
            testCase.mockBehavior(adv,testCase.inputAdvert)

            services:=&service.Service{Advert:adv}
            handler:=NewHandler(services)

            //Test server
            r:=gin.New()
            r.POST("/create_adv",handler.createAdvert)

            //Test request
            w:=httptest.NewRecorder()
            req:=httptest.NewRequest("POST","/create_adv",
                bytes.NewBufferString(testCase.inputBody))

            r.ServeHTTP(w,req)

            assert.Equal(t,testCase.expectedStatusCode,w.Code)
            assert.Equal(t,testCase.expectedRequestBody,w.Body.String())
        })
    }
}

func TestHandler_getAdvertById(t *testing.T){
    type mockBehavior func(s *mock_service.MockAdvert,id int, fields []string)

    testTable:=[]struct{
        name string
        id int
        inputBody string
        inputAdvertFields AdvertFields
        mockBehavior mockBehavior
        expectedStatusCode int
        expectedRequestBody string
    }{
        {
            name:"OK",
            id:1,
            inputBody:`{"fields":["description"]}`,
            inputAdvertFields: AdvertFields{
                Fields:[]string{"description"},
            },
            mockBehavior: func(s *mock_service.MockAdvert, id int, fields []string){
                s.EXPECT().GetAdvertById(id,fields).Return(map[string]interface{}{"title":"Title","description":"Description"},nil)
            },
            expectedStatusCode:200,
            expectedRequestBody:`{200,"status":200}`,
        },
    }

    for _,testCase:=range testTable{
        t.Run(testCase.name,func(t *testing.T){
            c:=gomock.NewController(t)
            defer c.Finish()

            adv:=mock_service.NewMockAdvert(c)
            testCase.mockBehavior(adv,testCase.id,testCase.inputAdvertFields.Fields)

            services:=&service.Service{Advert:adv}
            handler:=NewHandler(services)

            //Test server
            r:=gin.New()
            r.GET("/advert/:id",handler.getAdvertById)

            //Test request
            w:=httptest.NewRecorder()
            req:=httptest.NewRequest("GET","/advert/1",nil)

            r.ServeHTTP(w,req)

            assert.Equal(t,testCase.expectedStatusCode,w.Code)
            assert.Equal(t,testCase.expectedRequestBody,w.Body.String())
        })
    }
}