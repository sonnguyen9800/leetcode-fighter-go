package encode_decode_string

import (
	"strconv"
)

func encode_string(input_array_string []string) string {
	encoded_string := ""
	for _, str := range input_array_string {
		string_length := len(str)

		encoded_string += strconv.Itoa(string_length) + "#" + str
	}
	return encoded_string
}

func decode_string(input_string string) []string {
	original_string := []string{}

	current_string := ""
	string_length := 0
	string_counter := ""

	for _, char := range input_string {
		if string_length == 0 {

			string_counter += string(char)
			if char == '#' {
				string_counter = string_counter[:len(string_counter)-1]
				if val, err := strconv.Atoi(string_counter); err != nil {
					print(err)
				} else {
					string_length += val
					string_counter = "" // refresh
					if string_length == 0 {
						original_string = append(original_string, "")

					}
				}

			}
		} else {

			current_string += string(char)
			string_length--

			if string_length == 0 {
				original_string = append(original_string, current_string)
				current_string = "" //refresh
			}
			continue
		}

	}
	return original_string
}
