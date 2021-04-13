/*
Read Console Vars and save to a map

Example:

When you run

	go run *.go key1=value1 key2=value2

And need the value of a "key1":

	value1, found := mapconsolevars.GetValue("key1")
	if found {
		fmt.Println("Value of key1 is", value1)
	}

You can change the delimiter character with mapconsolevars.ChangeSeparatorTo.

Example:

	go run *.go key1:value1 key2:value2

And need the value of a "key1":

	mapconsolevars.ChangeSeparatorTo(":")
	value1, found := mapconsolevars.GetValue("key1")
	if found {
		fmt.Println("Value of key1 is", value1)
	}

*/
package mapconsolevars
