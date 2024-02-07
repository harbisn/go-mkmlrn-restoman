# go-mkmlrn-restoman

## postman curl
1. Create Menu
curl --location 'localhost:x/restoman/menu' \
--header 'x-user-id: 200275' \
--header 'Content-Type: application/json' \
--data '{
"name": "Shoyu Ramen",
"status": "AVAILABLE",
"category": "PASTA_AND_NODDLES",
"price": 50000,
"description": "Shoyu ramen is a ramen dish with a broth made of soy sauce."
}'

2. Get Menu with Pagination and Filter
curl --location 'localhost:x/restoman/menu?page=1&size=15&order=price%20DESC&status=NOT_AVAILABLE&lowestPrice=13000&category=MAIN_COURSE&highestPrice=25000'

3. Get Menu By Id
   curl --location 'localhost:x/restoman/menu/2'

4. Update Menu
   curl --location --request PATCH 'localhost:x/restoman/menu/2' \
   --header 'Content-Type: application/json' \
   --data '{
   "name": "Shoyu Ramen",
   "status": "NOT_AVAILABLE",
   "category": "PASTA_AND_NODDLES",
   "price": 40000,
   "description": "Shoyu ramen is a popular Japanese noodle balabalabala."
   }'

5. Delete Menu
   curl --location --request DELETE 'localhost:8080/restoman/menu/2' \
   --data ''

## Query insert to table
INSERT INTO menus ("name",status,category,price,description,created_by,created_at,updated_by,updated_at) VALUES
('Sushi','AVAILABLE','SUSHI',15000,'Assorted sushi rolls with fresh fish and rice.','200275','2024-02-07 14:05:00.000','200275','2024-02-07 14:05:00.000'),
('Tempura','AVAILABLE','APPETIZERS',12000,'Lightly battered and fried shrimp and vegetables.','200275','2024-02-07 14:10:00.000','200275','2024-02-07 14:10:00.000'),
('Teriyaki Chicken','AVAILABLE','MAIN_COURSE',18000,'Grilled chicken marinated in teriyaki sauce, served with steamed rice and vegetables.','200275','2024-02-07 14:15:00.000','200275','2024-02-07 14:15:00.000'),
('Miso Soup','AVAILABLE','SOUPS',5000,'Traditional Japanese soup made with miso paste, tofu, seaweed, and green onions.','200275','2024-02-07 14:20:00.000','200275','2024-02-07 14:20:00.000'),
('Yakitori','AVAILABLE','GRILLED',10000,'Skewered and grilled chicken with a savory sauce.','200275','2024-02-07 14:25:00.000','200275','2024-02-07 14:25:00.000'),
('Ramen','AVAILABLE','NOODLES',12000,'Japanese noodle soup with broth, noodles, meat, and vegetables.','200275','2024-02-07 14:30:00.000','200275','2024-02-07 14:30:00.000'),
('Okonomiyaki','AVAILABLE','PANCAKES',14000,'Japanese savory pancake made with cabbage, meat or seafood, and topped with okonomiyaki sauce and mayonnaise.','200275','2024-02-07 14:35:00.000','200275','2024-02-07 14:35:00.000'),
('Gyoza','AVAILABLE','APPETIZERS',8000,'Pan-fried Japanese dumplings filled with ground meat and vegetables.','200275','2024-02-07 14:40:00.000','200275','2024-02-07 14:40:00.000'),
('Udon','AVAILABLE','NOODLES',10000,'Thick wheat flour noodles served in a hot broth, often with toppings like tempura or tofu.','200275','2024-02-07 14:45:00.000','200275','2024-02-07 14:45:00.000'),
('Tonkatsu','AVAILABLE','MAIN_COURSE',16000,'Breaded and deep-fried pork cutlet served with shredded cabbage and tonkatsu sauce.','200275','2024-02-07 14:50:00.000','200275','2024-02-07 14:50:00.000');
INSERT INTO menus ("name",status,category,price,description,created_by,created_at,updated_by,updated_at) VALUES
('Chirashi Bowl','AVAILABLE','BOWLS',20000,'Assorted sashimi served over a bed of sushi rice.','200275','2024-02-07 14:55:00.000','200275','2024-02-07 14:55:00.000'),
('Sashimi Platter','AVAILABLE','PLATTERS',25000,'Variety of thinly sliced raw fish served with soy sauce and wasabi.','200275','2024-02-07 15:00:00.000','200275','2024-02-07 15:00:00.000'),
('Takoyaki','AVAILABLE','APPETIZERS',9000,'Japanese snack made of octopus pieces in a wheat-flour batter, cooked in a special molded pan.','200275','2024-02-07 15:05:00.000','200275','2024-02-07 15:05:00.000'),
('Onigiri','AVAILABLE','SNACKS',6000,'Japanese rice ball filled with various ingredients, wrapped in seaweed.','200275','2024-02-07 15:10:00.000','200275','2024-02-07 15:10:00.000'),
('Yakiniku','AVAILABLE','GRILLED',18000,'Japanese grilled meat, typically beef, cooked on a barbecue grill.','200275','2024-02-07 15:15:00.000','200275','2024-02-07 15:15:00.000'),
('Zaru Soba','AVAILABLE','NOODLES',11000,'Cold buckwheat noodles served with a dipping sauce and toppings like green onions and wasabi.','200275','2024-02-07 15:20:00.000','200275','2024-02-07 15:20:00.000'),
('Nikujaga','AVAILABLE','MAIN_COURSE',14000,'Japanese comfort food stew made with beef, potatoes, onions, and carrots simmered in a sweet soy sauce.','200275','2024-02-07 15:25:00.000','200275','2024-02-07 15:25:00.000'),
('Chawanmushi','AVAILABLE','APPETIZERS',7000,'Japanese savory egg custard dish with shrimp, mushrooms, and vegetables.','200275','2024-02-07 15:30:00.000','200275','2024-02-07 15:30:00.000'),
('Oyakodon','AVAILABLE','BOWLS',16000,'Japanese rice bowl dish topped with chicken, egg, and scallions cooked in a sweet and savory sauce.','200275','2024-02-07 15:35:00.000','200275','2024-02-07 15:35:00.000'),
('Matcha Ice Cream','AVAILABLE','DESSERTS',8000,'Japanese green tea flavored ice cream.','200275','2024-02-07 15:40:00.000','200275','2024-02-07 15:40:00.000');
INSERT INTO menus ("name",status,category,price,description,created_by,created_at,updated_by,updated_at) VALUES
('Tofu Teriyaki','AVAILABLE','MAIN_COURSE',14000,'Grilled tofu marinated in teriyaki sauce, served with steamed rice and vegetables.','200275','2024-02-07 15:45:00.000','200275','2024-02-07 15:45:00.000'),
('Yasai Tempura','AVAILABLE','APPETIZERS',10000,'Assorted lightly battered and fried vegetables.','200275','2024-02-07 15:50:00.000','200275','2024-02-07 15:50:00.000'),
('Chashu Ramen','AVAILABLE','NOODLES',13000,'Japanese noodle soup with broth, noodles, sliced pork belly, and vegetables.','200275','2024-02-07 15:55:00.000','200275','2024-02-07 15:55:00.000'),
('Yuzu Shoyu Ramen','AVAILABLE','NOODLES',15000,'Shoyu ramen with a citrus twist from yuzu, served with sliced pork and green onions.','200275','2024-02-07 16:00:00.000','200275','2024-02-07 16:00:00.000'),
('Ankake Yakisoba','AVAILABLE','NOODLES',11000,'Stir-fried noodles with vegetables and meat, topped with a thick savory sauce.','200275','2024-02-07 16:05:00.000','200275','2024-02-07 16:05:00.000'),
('Agedashi Tofu','AVAILABLE','APPETIZERS',9000,'Deep-fried tofu served in a flavorful broth with grated daikon, ginger, and green onions.','200275','2024-02-07 16:10:00.000','200275','2024-02-07 16:10:00.000'),
('Katsu Curry','AVAILABLE','MAIN_COURSE',17000,'Breaded and deep-fried meat (chicken, pork, or beef) served with Japanese curry sauce and rice.','200275','2024-02-07 16:15:00.000','200275','2024-02-07 16:15:00.000'),
('Sukiyaki','AVAILABLE','MAIN_COURSE',19000,'Japanese hot pot dish with thinly sliced beef, vegetables, tofu, and noodles cooked in a sweet soy sauce broth.','200275','2024-02-07 16:20:00.000','200275','2024-02-07 16:20:00.000'),
('Tuna Tataki','AVAILABLE','APPETIZERS',16000,'Seared slices of fresh tuna served with ponzu sauce, grated ginger, and garlic.','200275','2024-02-07 16:25:00.000','200275','2024-02-07 16:25:00.000'),
('Matcha Cheesecake','AVAILABLE','DESSERTS',12000,'Creamy cheesecake infused with matcha green tea flavor.','200275','2024-02-07 16:30:00.000','200275','2024-02-07 16:30:00.000');
