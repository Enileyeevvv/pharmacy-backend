package postgres

const (
	queryFetchMedicines = `
		select mp.id as id,
			   mp.name as name,
			   mp.sell_name as sell_name,
			   mp.atx_code as atx_code,
			   mp.description as description,
			   mp.pharmaceutical_group_id as pharmaceutical_group_id,
			   pg.name as pharmaceutical_group_name,
			   mp.quantity as quantity,
			   mp.max_quantity as max_quantity,
			   c.id as company_id,
			   c.name as company_name,
			   mpc.image_url as image_url
	
		from medicinal_product_company mpc
			join medicinal_products mp
				on mp.id = mpc.medicinal_product_id
			join company c
				on mpc.company_id = c.id
			join pharmaceutical_group pg
				on mp.pharmaceutical_group_id = pg.id
		order by id
		limit ($1 + 1) offset ($1 * ($2 - 1));
`
)
