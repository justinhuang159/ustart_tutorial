package elasticstore

import "context"

func (estor *ElasticStore) LookupCarAvailiabilty(ctx context.Context, request string) (bool, error){
	return estor.Lookup(ctx, request)
}