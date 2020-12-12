package action

import (
	"github.com/ritarock/athenaq/lib/aws"
)

func Run(profile, bucket, key, query string) {
	sess := aws.Session(profile)
	aws.RunAthenaQuery(sess, bucket, key, query)
}
