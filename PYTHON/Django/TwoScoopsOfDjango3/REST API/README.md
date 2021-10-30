# 17 Building REST APIs With Django REST Framework

## 1 Fundamentals ofr Basic REST API Design
*   _GET_, _PUT_ and _DELETE_ are idempotent. _POST_ and _PATCH_ are not
*   _PATCH_ is ofter not implemented, but it's good idea to implement it if your API supports _PUT_


## 3 REST API 
*   Use Consistent API Module Naming
* Code for a Project Should Be Neatly Organized
* Code for an App Should Remain in the App
*   Try to keep business Logic Out of API Views
* Test Your API
* Version your API(maintain both existing and predecessor API, If API new version come or deprecated Warn Costumers, you can use [django rest](https://www.django-rest-framework.org/api-guide/versioning/) default versioning system)
*   Be careful with customized Authentication Schemes
    * Also Document your Authentication System force you to have deeper understanding over your system


## 5 Shutdown and external API

* Step #1: Notify Users of Pending Shutdown
* Step #1: Replace API with 410 Error View

## 6 Rete-Limiting Your API

## 7 Advertising Your REST API
* Documentation 
    *   Write down from scratch 
    * or Use third party application for example
        * Open API
        * Swagger
* Provide SDKs