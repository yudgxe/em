package enricher

import (
	"context"
	"testing"
)

func Test_Enricher(t *testing.T) {
	t.Skip()

	const expectedAge int32 = 62

	var enricher Enricher
	age, err := enricher.GetAgeByName(context.Background(), "test")
	if err != nil {
		t.Fatal(err)
	}

	if got, want := *age, expectedAge; got != want {
		t.Fatalf("got: %v, want: %v", got, want)
	}

	nat, err := enricher.GetNationalityByName(context.Background(), "test")
	if err != nil {
		t.Fatal(err)
	}

	t.Fatal(nat)
}
