package elasticStore

import "context"

func (estor *ElasticStore) UpdateCarAvailability(ctx context.Context, CID string, status bool) (error){
	_, err = elasticClient.Update().
		Index(estor.eIndex).
		Type(estor.eType).
		Id(CID).
		Doc(map[string]interface{}{"Available": status}).
		Do(ctx)
	return err
}
		