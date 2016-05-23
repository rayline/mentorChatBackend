#Database document of backend of mentorChat

##Summary
	We now use redis as the database since we are not running a application with big load.

##Points of database use

+ storing user basic information
+ implemention of MESSAGE queue
+ friend list maintaining
+ UID allocating
+ mail to uid
+ name to uid

##The way we do things

###basics
+ representing numbers in decmical

###storing user basic information
+ database 0
+ type: string
+ encoding: json

###implemention of MESSAGE queue
+ database 1
+ type: list
+ encoding: json

###friend list maintaining
+ database 2
+ type: set
+ encoding: dec for uid

###UID allocating
+ database 0
+ type: string
+ name: NEXTUID
+ default: "1"