Coupon API
==========

Table of Contents
-----------------

*   [Overview](#overview)
    
*   [Implemented Use Cases](#implemented-use-cases)
    
*   [Unimplemented Use Cases](#unimplemented-use-cases)
    
*   [Technical Documentation](#technical-documentation)
    
    *   [Project Structure](#project-structure)
        
    *   [Data Models](#data-models)
        
    *   [Architecture](#architecture)
        
    *   [Dependency Injection](#dependency-injection)
        
*   [Instructions to Run the Project](#instructions-to-run-the-project)
    
    *   [Prerequisites](#prerequisites)
        
    *   [Steps](#steps)
        
*   [API Documentation (Swagger)](#api-documentation-swagger)
    
    *   [Installing Swag](#installing-swag)
        
    *   [Adding Swagger Annotations](#adding-swagger-annotations)
        
    *   [Generating Swagger Documentation](#generating-swagger-documentation)
        
    *   [Accessing Swagger UI](#accessing-swagger-ui)
        
*   [Examples](#examples)
    
    *   [Creating Coupons](#creating-coupons)
        
    *   [Fetching Applicable Coupons](#fetching-applicable-coupons)
        
    *   [Applying Coupons](#applying-coupons)
        
*   [Sample cURL Commands](#sample-curl-commands)
    
    *   [Create a Coupon](#create-a-coupon)
        
    *   [Get All Coupons](#get-all-coupons)
        
    *   [Apply a Coupon](#apply-a-coupon)
        
*   [Assumptions and Limitations](#assumptions-and-limitations)
    
    *   [Assumptions](#assumptions)
        
    *   [Limitations](#limitations)
        
*   [Suggestions for Improvement](#suggestions-for-improvement)
    
*   [Conclusion](#conclusion)
    
*   [Contact](#contact)
    

Overview
--------

This project is a **RESTful API** for managing and applying different types of discount coupons for an e-commerce platform. It is designed to be **extensible**, allowing for easy addition of new coupon types in the future.

The API supports:

*   **Cart-wise Coupons**: Apply discounts to the entire cart based on certain conditions.
    
*   **Product-wise Coupons**: Apply discounts to specific products.
    
*   **BxGy Coupons**: "Buy X, Get Y" deals with repetition limits and applicable to sets of products.
    
*   **Limited-use Coupons**: Coupons that can only be used a certain number of times globally.
    
*   **User-specific Coupons**: Coupons assigned to specific users (structure in place, logic not fully implemented).
    

Implemented Use Cases
---------------------

1.  **Cart-wise Coupons**
    
    *   **Percentage Discount over Threshold**
        
        *   **Condition**: Cart total exceeds a specified threshold.
            
        *   **Discount**: A percentage discount on the entire cart total.
            
    *   **Fixed Amount Discount over Threshold**
        
        *   **Condition**: Cart total exceeds a specified threshold.
            
        *   **Discount**: A fixed amount discount on the entire cart total.
            
    *   **Tiered Discounts**
        
        *   **Condition**: Different discount rates based on cart total tiers.
            
        *   **Discount**: Percentage off based on cart total.
            
2.  **Product-wise Coupons**
    
    *   **Percentage Discount on Specific Product**
        
        *   **Condition**: Specific product is in the cart.
            
        *   **Discount**: Percentage off on that product.
            
    *   **Fixed Discount on Specific Product**
        
        *   **Condition**: Specific product is in the cart.
            
        *   **Discount**: Fixed amount off on that product.
            
3.  **BxGy Coupons**
    
    *   **Simple BxGy**
        
        *   **Condition**: Buy X quantity of a product.
            
        *   **Discount**: Get Y quantity of the same product free.
            
    *   **Cross Product BxGy**
        
        *   **Condition**: Buy X quantity of Product A.
            
        *   **Discount**: Get Y quantity of Product B free.
            
    *   **Mixed Products BxGy**
        
        *   **Condition**: Buy X products from an array of products.
            
        *   **Discount**: Get Y products from another array of products free.
            
    *   **Repetition Limit**
        
        *   **Condition**: Coupon can be applied up to a repetition limit.
            
        *   **Discount**: Applied multiple times based on the limit.
            
4.  **Limited-use Coupons**
    
    *   **Usage Limits**
        
        *   **Condition**: Coupon can only be used a certain number of times globally.
            
        *   **Implementation**: Tracks coupon usage count.
            

Unimplemented Use Cases
-----------------------

1.  **User-specific Coupons**
    
    *   Coupons assigned to specific users.
        
    *   Requires user authentication and authorization.
        
2.  **Time-based Coupons**
    
    *   Coupons valid only during specific hours or days.
        
    *   Requires time-based checks and potentially a scheduler.
        
3.  **First-time Buyer Coupons**
    
    *   Coupons applicable only on a user's first purchase.
        
    *   Requires tracking user purchase history.
        
4.  **Referral Coupons**
    
    *   Coupons provided when a user refers another user.
        
    *   Involves user relationships and referral tracking.
        
5.  **Category-wise Discounts**
    
    *   Apply discounts to entire product categories.
        
    *   Requires product categorization.
        
6.  **Coupon Stacking**
    
    *   Allowing multiple coupons to be applied to a single cart.
        
    *   Requires defining stacking rules and resolving conflicts.
        
7.  **Loyalty Program Integration**
    
    *   Discounts based on loyalty points or membership levels.
        
    *   Requires a loyalty program system.
        


### Architecture

*   **Models**: Define the data structures used in the application.
    
*   **Repositories**: Handle data storage and retrieval. In this case, data is stored in a JSON file (data/coupons.json).
    
*   **Services**: Contain business logic and interact with repositories and strategies.
    
*   **Strategies**: Implement the logic for different coupon types.
    
*   **Handlers**: Handle HTTP requests and responses using Gin Gonic.
    
*   **Data Storage**: Uses a JSON file for simplicity and ease of setup.
    

Instructions to Run the Project
-------------------------------

### Prerequisites

*   **Go**: Version 1.16 or higher installed on your machine.
    
*   **Git**: For cloning the repository.
    

### Steps

1.  Copy code[git clone https://github.com/rungta-ayush/monk.git](https://github.com/rungta-ayush/monk)**cd coupon-api**
    
2.  Run the following command to download the required Go modules:
    

go mod download

1.  Ensure the data directory exists and contains a coupons.json file. If it doesn't exist, create it: mkdir -p dataecho "\[\]" > data/coupons.json
    
2.  Start the server using the following command:go run main.goThe server will start on http://localhost:8080.
    
3.  Use tools like **Postman** or **cURL** to interact with the API endpoints.
    
4.  Follow the instructions in the [API Documentation (Swagger)](#api-documentation-swagger) section to generate and access the Swagger UI.
    

API Documentation (Swagger)
---------------------------

Access the Swagger UI at http://localhost:8080/swagger

Examples
--------

### Creating Coupons

#### Example 1: Creating a Cart-wise Coupon

**Request**

`   POST /coupons  Content-Type: application/json  {    "type": "cart-wise",    "details": {      "threshold": 100,      "discount": 10    },    "expiration_date": null  }   `

**Response**

`   {    "id": 1,    "type": "cart-wise",    "details": {      "threshold": 100,      "discount": 10    },    "expiration_date": null,    "usage_limit": 0,    "used_count": 0,    "users": []  }   `

#### Example 2: Creating a BxGy Coupon

**Request**

`   POST /coupons  Content-Type: application/json  {    "type": "bxgy",    "details": {      "buy_products": [        { "product_id": 1, "quantity": 2 },        { "product_id": 2, "quantity": 1 }      ],      "get_products": [        { "product_id": 3, "quantity": 1 }      ],      "repetition_limit": 3    },    "expiration_date": null  }   `

**Response**

`   {    "id": 2,    "type": "bxgy",    "details": {      "buy_products": [        { "product_id": 1, "quantity": 2 },        { "product_id": 2, "quantity": 1 }      ],      "get_products": [        { "product_id": 3, "quantity": 1 }      ],      "repetition_limit": 3    },    "expiration_date": null,    "usage_limit": 0,    "used_count": 0,    "users": []  }   `

### Fetching Applicable Coupons

**Request**

`   POST /applicable-coupons  Content-Type: application/json  {    "cart": {      "user_id": 123,      "items": [        { "product_id": 1, "quantity": 4, "price": 50 },        { "product_id": 2, "quantity": 2, "price": 30 },        { "product_id": 3, "quantity": 1, "price": 25 }      ]    }  }   `

**Response**

`   {    "applicable_coupons": [      {        "coupon_id": 1,        "type": "cart-wise",        "discount": 35      },      {        "coupon_id": 2,        "type": "bxgy",        "discount": 25      }    ]  }   `

### Applying Coupons

**Request**

Plain textANTLR4BashCC#CSSCoffeeScriptCMakeDartDjangoDockerEJSErlangGitGoGraphQLGroovyHTMLJavaJavaScriptJSONJSXKotlinLaTeXLessLuaMakefileMarkdownMATLABMarkupObjective-CPerlPHPPowerShell.propertiesProtocol BuffersPythonRRubySass (Sass)Sass (Scss)SchemeSQLShellSwiftSVGTSXTypeScriptWebAssemblyYAMLXML`   POST /apply-coupon/2  Content-Type: application/json  {    "cart": {      "user_id": 123,      "items": [        { "product_id": 1, "quantity": 4, "price": 50 },        { "product_id": 2, "quantity": 2, "price": 30 },        { "product_id": 3, "quantity": 1, "price": 25 }      ]    }  }   `

**Response**
`   {    "updated_cart": {      "items": [        { "product_id": 1, "quantity": 4, "price": 50, "total_discount": 0 },        { "product_id": 2, "quantity": 2, "price": 30, "total_discount": 0 },        { "product_id": 3, "quantity": 1, "price": 25, "total_discount": 25 }      ],      "total_price": 285,      "total_discount": 25,      "final_price": 260    }  }   `

Sample cURL Commands
--------------------

Below are sample cURL commands to interact with the API.

### Create a Coupon

`   curl -X POST http://localhost:8080/coupons \  -H 'Content-Type: application/json' \  -d '{    "type": "cart-wise",    "details": {      "threshold": 100,      "discount": 10    },    "expiration_date": null  }'   `

### Get All Coupons

`   curl http://localhost:8080/coupons   `

### Apply a Coupon

`   curl -X POST http://localhost:8080/apply-coupon/1 \  -H 'Content-Type: application/json' \  -d '{    "cart": {      "user_id": 123,      "items": [        { "product_id": 1, "quantity": 2, "price": 50 },        { "product_id": 2, "quantity": 1, "price": 30 }      ]    }  }'   `

Assumptions and Limitations
---------------------------

### Assumptions

*   **Single Coupon Application**: Only one coupon can be applied per cart.
    
*   **Price Stability**: Product prices do not change during the application process.
    
*   **Time Zone**: All times are in UTC.
    
*   **Data Integrity**: Coupons stored in the JSON file are correctly formatted.
    
*   **User Authentication**: Not implemented; user IDs are assumed to be valid.
    

### Limitations

*   **Concurrency**: Basic mutex locking is used; may not scale well with high concurrency.
    
*   **Error Handling**: Error messages are informative but could be expanded with error codes.
    
*   **Persistence**: Using a JSON file for data storage is not suitable for production environments.
    
*   **Data Validation**: Basic validation is performed, but more extensive validation could be added.
    
*   **Unimplemented Features**: Some coupon types and use cases are not fully implemented.
    

Suggestions for Improvement
---------------------------

*   **User Authentication and Authorization**: Implement user-specific coupons and secure endpoints.
    
*   **Database Integration**: Use a proper database (e.g., PostgreSQL, MongoDB) for data storage.
    
*   **Enhanced Validation**: Add more thorough validation and error handling.
    
*   **Logging and Monitoring**: Integrate logging for better debugging and monitoring.
    
*   **Scalability**: Optimize the application for high load and concurrency.
    
*   **Coupon Stacking**: Allow multiple coupons to be applied simultaneously with defined stacking rules.
    
*   **Category-wise Discounts**: Implement discounts applicable to entire product categories.
    
*   **Time-based Coupons**: Fully implement time-based coupon logic.
    
*   **API Versioning**: Introduce versioning for the API endpoints.
    
*   **Testing**: Add unit tests and integration tests for better reliability.