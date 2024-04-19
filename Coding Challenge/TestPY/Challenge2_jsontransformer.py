#!/usr/bin/env python
# coding: utf-8

# In[33]:


import json
from datetime import datetime

def transform_json(input_data):
    output = []

    for key, value in input_data.items():
        key = key.strip()  # Sanitize the key
        
        # Skip empty keys and invalid fields
        if not key or isinstance(value, dict) and not value:
            continue
        
        # Transform value based on data type
        if isinstance(value, dict):
            data_type = list(value.keys())[0]
            actual_value = list(value.values())[0].strip()  # Sanitize the value

            if data_type == "S":
                # Transform RFC3339 formatted strings to Unix Epoch
                try:
                    actual_value = int(datetime.fromisoformat(actual_value.replace('Z', '+00:00')).timestamp())
                except ValueError:
                    pass  # If not RFC3339 formatted, keep as is
                
                if actual_value == "":
                    continue  # Omit fields with empty values
                
                output.append({key: actual_value})

            elif data_type == "N":
                try:
                    actual_value = int(actual_value.lstrip('0')) if '.' not in actual_value else float(actual_value)
                    output.append({key: actual_value})
                except ValueError:
                    pass  # Skip invalid number values
            
            elif data_type == "BOOL":
                if actual_value.lower() in ["1", "t", "true"]:
                    actual_value = True
                elif actual_value.lower() in ["0", "f", "false"]:
                    actual_value = False
                else:
                    continue  # Skip invalid BOOL values
                output.append({key: actual_value})

            elif data_type == "NULL":
                if actual_value.lower() in ["1", "t", "true"]:
                    actual_value = None
                elif actual_value.lower() in ["0", "f", "false"]:
                    continue  # Omit fields with null value false
                output.append({key: actual_value})

            elif data_type == "L":
                # Transform list values
                list_values = [item for item in actual_value.split(',') if item.strip()]
                transformed_list = []
                for item in list_values:
                    if item.isdigit():
                        transformed_list.append(int(item.lstrip('0')))
                    elif item.lower() in ["1", "t", "true"]:
                        transformed_list.append(True)
                    elif item.lower() in ["0", "f", "false"]:
                        transformed_list.append(False)
                if transformed_list:
                    output.append({key: transformed_list})

            elif data_type == "M":
                # Recursively transform nested map
                transformed_map = transform_json(actual_value)
                if transformed_map:
                    output.append({key: transformed_map})

    return output

# Input JSON data
input_json = {
    "number_1": {"N": "1.50"},
    "string_1": {"S": "784498 "},
    "string_2": {"S": "2014-07-16T20:55:46Z"},
    "map_1": {
        "M": {
            "bool_1": {"BOOL": "truthy"},
            "null_1": {"NULL ": "true"},
            "list_1": {"L": ",011,5215s,f,0"}
        }
    },
    "list_2": {"L": "noop"},
    "list_3": {"L": ["noop"]},
    "": {"S": "noop"}
}

# Transform JSON and print to stdout
output_json = transform_json(input_json)
print(json.dumps(output_json, indent=2))


# In[ ]:


package main

import (
	"encoding/json"
	"fmt"
	"os"
	"reflect"
	"strconv"
	"strings"
)

func main() {
	// Declaration
	var input map[string]interface{}

	// Reading file
	data, err := os.ReadFile("input.json")
	if err != nil {
		fmt.Println(err)
	}

	// JSON conversion
	json.Unmarshal(data, &input)

	// Performing Transition Operation
	output := constructMap(input)

	// Priting output in human readable format
	jsonString, err := json.MarshalIndent(output, "", "    ")
	if err != nil {
		fmt.Println("Error marshalling JSON:", err)
		return
	}

	// Print the formatted JSON string
	fmt.Println(string(jsonString))
}

func constructMap(input map[string]interface{}) map[string]interface{} {
	// Reading Map data type

	output := make(map[string]interface{})

	for pk, pv := range input {
		if pk == "" {
			continue
		}
		for k, v := range pv.(map[string]interface{}) {
			if k == "" || v == "" {
				continue
			}

			switch k {
			case "N":
				status, rtnData := validateType("number", v)

				if status {
					output[pk] = rtnData
				}

			case "S":
				status, rtnData := validateType("string", v)

				if status {
					output[pk] = rtnData
				}
			case "M":
				data, ok := v.(map[string]interface{})

				if !ok {
					output[pk] = v
				} else {
					output[pk] = constructMap(data)
				}

			case "L":
				rtnData := constructList(v)

				if len(rtnData) > 0 {
					output[pk] = rtnData
				}
			case "BOOL":
				status, rtnData := validateType("bool", v)

				if status {
					output[pk] = rtnData
				}

			case "NULL":
				status, rtnData := validateType("null", v)

				if status {
					output[pk] = rtnData
				}

			default:
				output[pk] = v
			}
		}

	}

	return output

}

func constructList(data interface{}) []interface{} {
	// Reading list of array data type

	var out []interface{}

	switch data.(type) {
	case []interface{}:

		for _, pv := range data.([]interface{}) {

			if reflect.TypeOf(pv).Kind() == reflect.Map {
				for k, v := range pv.(map[string]interface{}) {
					if k == "" || v == "" {
						continue
					}

					switch k {
					case "S":
						status, rtnData := validateType("string", v)

						if status {
							out = append(out, rtnData)
						}
					case "N":
						status, rtnData := validateType("number", v)

						if status {
							out = append(out, rtnData)
						}
					case "BOOL":
						status, rtnData := validateType("bool", v)

						if status {
							out = append(out, rtnData)
						}
					case "NULL":

						status, rtnData := validateType("null", v)

						if status {
							out = append(out, rtnData)
						}
					}
				}
			}
		}

	}

	return out
}

func validateType(dataType string, val interface{}) (bool, interface{}) {

	// Utility function to validate and convert the type

	switch dataType {
	case "string":
		cnvData, ok := val.(string)
		if !ok {
			return false, val
		}

		return true, strings.Trim(cnvData, " ")

	case "number":
		strValue := val.(string)

		// Convert the string to a float64
		cnvData, err := strconv.ParseFloat(strValue, 64)

		if err != nil {
			return false, val
		}

		return true, cnvData
	case "bool":
		if strings.ToLower(val.(string)) == "true" || strings.ToLower(val.(string)) == "t" || strings.ToLower(val.(string)) == "1" {
			return true, true
		} else if strings.ToLower(val.(string)) == "false" || strings.ToLower(val.(string)) == "f" || strings.ToLower(val.(string)) == "0" {
			return true, false
		}

		return false, val

	case "null":
		return val.(string) == "", val.(string)
	default:
		return false, val
	}

}

