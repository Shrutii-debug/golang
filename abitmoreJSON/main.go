package main

import (
	"encoding/json"
	"fmt"
)

//using lowercase c as we dont have to export it anywhere or make it public
type course struct {
	Name  string `json:"coursename"`
	Price int		
	Platform string	`json:"website"`
	Password string	`json:"-"` //this dash means that i dont want this field to be reflected whoever is using my API
	Tags  []string    `json:"tags,omitempty"`        //slice of strings
}


func main() {
	fmt.Println("Welcome to a bit more about JSON")

	 // EncodeJson()
	DecodeJson()
}

func EncodeJson(){

	lcoCourses := []course{
		{"ReactJS Bootcamp", 299, "LearnCodeOnline.in", "abc123", []string{"web-dev", "js"}},
		{"MERN Bootcamp", 199, "LearnCodeOnline.in", "bcd123", []string{"full-stack", "js"}},
		{"Angular Bootcamp", 299, "LearnCodeOnline.in", "shr123", nil},
	}

	//package this data as JSON data

	//Marshal is an implementation of JSON, its an alternative version of struct
	//finalJson, err := json.Marshal(lcoCourses)
	finalJson, err := json.MarshalIndent(lcoCourses, "", "\t") //this takes two parameters, first is the interface or struct itself, 
	//and the second is based on what should I indent the values because it doesnt know youre throwing up the JSON, it can be used for other things as well
	if err != nil{
		panic(err)
	}
	fmt.Printf("%s\n", finalJson)
}

//IMPORTANT
/*
json.MarshalIndent in Go is used when you want pretty-printed (formatted) JSON instead of the compact one produced by json.Marshal.

It is very useful when:

printing JSON for debugging

creating readable API responses

writing JSON files

1️⃣ Syntax
json.MarshalIndent(v interface{}, prefix string, indent string)
Parameters
Parameter	Meaning
v	the data you want to convert into JSON
prefix	string added at the beginning of every line
indent	indentation used for nested levels

Return values:

[]byte, error

Just like json.Marshal.

2️⃣ Example Without Indentation (json.Marshal)
package main

import (
	"encoding/json"
	"fmt"
)

type Course struct {
	Name     string
	Price    int
	Platform string
}

func main() {

	course := Course{"React Bootcamp", 299, "LearnCodeOnline"}

	data, _ := json.Marshal(course)

	fmt.Println(string(data))
}

Output (compact JSON):

{"Name":"React Bootcamp","Price":299,"Platform":"LearnCodeOnline"}

Everything is in one line.

3️⃣ Example Using json.MarshalIndent
package main

import (
	"encoding/json"
	"fmt"
)

type Course struct {
	Name     string
	Price    int
	Platform string
}

func main() {

	course := Course{"React Bootcamp", 299, "LearnCodeOnline"}

	data, _ := json.MarshalIndent(course, "", "  ")

	fmt.Println(string(data))
}

Output:

{
  "Name": "React Bootcamp",
  "Price": 299,
  "Platform": "LearnCodeOnline"
}

Now it is formatted and readable.

4️⃣ Understanding the Parameters
1️⃣ First parameter → Data
json.MarshalIndent(course, "", "  ")

Here course is the struct being converted to JSON.

2️⃣ Second parameter → Prefix

This is added before every line.

Example:

data, _ := json.MarshalIndent(course, ">> ", "  ")

Output:

>> {
>>   "Name": "React Bootcamp",
>>   "Price": 299,
>>   "Platform": "LearnCodeOnline"
>> }

Usually we keep it:

""
3️⃣ Third parameter → Indentation

This defines spacing for nested JSON levels.

Example:

data, _ := json.MarshalIndent(course, "", "\t")

Output:

{
	"Name": "React Bootcamp",
	"Price": 299,
	"Platform": "LearnCodeOnline"
}

Here indentation uses tabs instead of spaces.

5️⃣ Example With Slice (Similar to Your Code)
type Course struct {
	Name     string
	Price    int
	Platform string
}

func main() {

	courses := []Course{
		{"React Bootcamp", 299, "LearnCodeOnline"},
		{"MERN Bootcamp", 199, "LearnCodeOnline"},
	}

	data, _ := json.MarshalIndent(courses, "", "  ")

	fmt.Println(string(data))
}

Output:

[
  {
    "Name": "React Bootcamp",
    "Price": 299,
    "Platform": "LearnCodeOnline"
  },
  {
    "Name": "MERN Bootcamp",
    "Price": 199,
    "Platform": "LearnCodeOnline"
  }
]
6️⃣ Real Use Case (Very Common)

Writing formatted JSON to a file.

data, _ := json.MarshalIndent(courses, "", "  ")

os.WriteFile("courses.json", data, 0644)

This creates a clean readable JSON file.

7️⃣ Difference Between Marshal and MarshalIndent
Function	Output
json.Marshal	compact JSON
json.MarshalIndent	formatted JSON

Example:

Marshal
{"Name":"React","Price":299}
MarshalIndent
{
  "Name": "React",
  "Price": 299
}
8️⃣ Important Thing (Common Beginner Confusion)

Both return:

[]byte

So you must print like this:

fmt.Println(string(data))

or

fmt.Printf("%s\n", data)

✅ In short

json.Marshal        → compact JSON
json.MarshalIndent  → pretty JSON

Parameters:

(data, prefix, indentation)

Example:

json.MarshalIndent(data, "", "  ")

💡 Since you're learning Go JSON encoding, the next extremely important concept you’ll encounter is:

struct tags

Example:

Name string `json:"course_name"`
Password string `json:"-"`
Price int `json:"price,omitempty"`

These control how JSON is generated, and they are used in almost every Go backend API.
*/

//important
/*
he part inside backticks ` ` in your struct is called a struct tag in Go.
In your case, it is a JSON struct tag, which tells the encoding/json package how to convert your struct into JSON.

Let's go through your code step by step.

1️⃣ What are Struct Tags?

Struct tags are metadata attached to struct fields.

Syntax:

FieldName Type `key:"value"`

Example:

Name string `json:"coursename"`

This tells Go:

When converting this struct to JSON, use "coursename" instead of "Name".

2️⃣ Your Struct
type course struct {
	Name     string   `json:"coursename"`
	Price    int
	Platform string   `json:"website"`
	Password string   `json:"-"`
	Tags     []string `json:"tags,omitempty"`
}

Now let's explain each field.

3️⃣ Name string json:"coursename"`

Normally Go would produce:

{
  "Name": "ReactJS Bootcamp"
}

But because of the tag:

json:"coursename"

the JSON becomes:

{
  "coursename": "ReactJS Bootcamp"
}

So this tag renames the field in JSON.

4️⃣ Price int

Since there is no JSON tag, Go uses the default struct field name.

Output:

{
  "Price": 299
}
5️⃣ Platform string json:"website"`

The tag changes the JSON key.

Struct field:

Platform string

JSON output:

{
  "website": "LearnCodeOnline.in"
}
6️⃣ Password string json:"-"``

The dash - means:

Do not include this field in JSON at all.

Example struct:

Password: "abc123"

JSON output will NOT include it.

Result:

{
  "coursename": "ReactJS Bootcamp",
  "Price": 299,
  "website": "LearnCodeOnline.in"
}

Password is hidden.

This is commonly used for:

passwords

tokens

internal fields

7️⃣ Tags []string json:"tags,omitempty"`

Here two things happen.

1️⃣ tags

This sets the JSON key name.

Tags → tags
2️⃣ omitempty

This means:

If the field is empty, do not include it in JSON.

Empty means:

""

0

nil

[]

false

Example:

Case 1
Tags: []string{"web-dev", "js"}

JSON:

"tags": ["web-dev","js"]
Case 2
Tags: nil

JSON:

(no tags field at all)

Instead of:

"tags": null
8️⃣ Why Struct Fields Must Start with Capital Letters

You used:

type course struct

But fields:

Name
Price
Platform
Password
Tags

are capitalized.

This is required because:

Go's JSON package can only access exported fields.

Rule:

Field	JSON Works?
Name	✅
name	❌

If you write:

name string

it will not appear in JSON.

9️⃣ Final JSON Output of Your Struct

Your program will produce something like:

[
  {
    "coursename": "ReactJS Bootcamp",
    "Price": 299,
    "website": "LearnCodeOnline.in",
    "tags": ["web-dev","js"]
  },
  {
    "coursename": "MERN Bootcamp",
    "Price": 199,
    "website": "LearnCodeOnline.in",
    "tags": ["full-stack","js"]
  },
  {
    "coursename": "Angular Bootcamp",
    "Price": 299,
    "website": "LearnCodeOnline.in"
  }
]

Notice:

Password is missing (because of -)

Tags missing in last course (because of omitempty)

🔑 Quick Summary
Tag	Meaning
json:"name"	rename JSON field
json:"-"	exclude field
json:",omitempty"	hide if empty
json:"name,omitempty"	rename + hide if empty

💡 One more super important Go JSON concept (which you'll hit next in backend APIs):

json.Unmarshal()

This converts JSON → struct (reverse of Marshal).

Example:

JSON → Go struct → API logic
*/


//consuming the JSON data

func DecodeJson(){
	jsonDataFromWeb := []byte(`
		 {
                "coursename": "ReactJS Bootcamp",
                "Price": 299,
                "website": "LearnCodeOnline.in",
                "tags": ["web-dev","js"]
        }
	`)

	var lcoCourse course 

	checkValid := json.Valid(jsonDataFromWeb) //validates the json data and gives a boolean result

	if checkValid{
		fmt.Println("JSON was valid")
		json.Unmarshal(jsonDataFromWeb, &lcoCourse) //unmarshal is a method and we are passing on a reference of our interface(Struct)
		fmt.Printf("%#v\n", lcoCourse)
	} else {
		fmt.Println("JSON data not valid")
	}

	//some cases where you just want to add data to key value pairs

	var myOnlineData map[string]interface{}  //we dont know if the data that is coming, all of it is a string or anything,so we use interface
	json.Unmarshal(jsonDataFromWeb, &myOnlineData)
	fmt.Printf("%#v\n", myOnlineData)

	for k, v := range myOnlineData{
		fmt.Printf("key is %v and value is %v and type is: %T \n", k, v, v)
	}
}