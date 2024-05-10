package postgres

const (
	queryFetchMedicines = `
		select id,
			   name,
			   sell_name,
			   atx_code,
			   description,
			   pharmaceutical_group_id,
			   quantity,
			   max_quantity
		from medicinal_products
		order by id
		limit ($1 + 1) offset ($1 * ($2 - 1));
`
)
