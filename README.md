# urlStatusChecker


for using this api server 

1. GET request   
    a. For getting status of all webisites use          ( curl "http://localhost:8080/Websites" )
    b. For getting status of particular website use     ( curl "http://localhost:8080/Websites?name=www.google.com" )
    
2. POST request
   For adding the website in your server use 
                     curl http://localhost:8080/Websites \                  
           --header "Content-Type: application/json" \
              --request "POST" \
              --data '["www.google.com", "www.facebook.com", "www.fakebook123.com"]'

 
