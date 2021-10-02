package sonarscanner

type ProjectStatusCondition struct {
	Status	string	`json:status`
	Key		string	`json:metricKey`
	Value	string	`json:actualValue`
}