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

	queryCreateMedicalProduct = `
		insert into medicinal_products
		(name, sell_name, atx_code, description, quantity, max_quantity, pharmaceutical_group_id)
		values ($1, $2, $3, $4, $5, $6, $7)
		returning id;
`

	queryCheckMedicinalProductExists = `
		select id
		from medicinal_products
		where (trim(lower(name))) = $1
		  and (trim(lower(sell_name))) = $2;
`

	queryCheckCompanyExists = `
		select id
		from company
		where (trim(lower(name))) = $1;
`

	queryCreateCompany = `
		insert into company 
		    (name) 
		values ($1) 
		returning id;
`

	queryUpsertMedicinalProductCompany = `
		insert into medicinal_product_company
			(medicinal_product_id, company_id, image_url)
		values ($1, $2, $3)
		on conflict (medicinal_product_id, company_id) 
		    do update 
		    set image_url = excluded.image_url;
`

	queryFetchPatients = `
		select p.id as id,
			   p.name as name,
			   p.email as email,
			   p.birthday as birthday,
			   p.createdat as created_at,
			   p.updatedat as updated_at
		from patient p
		where ($3::text is null or name = $3::text)
		order by id
		limit ($1 + 1) offset ($1 * ($2 - 1));
`

	queryGetPatient = `
		select p.id as id,
			   p.name as name,
			   p.email as email,
			   p.birthday as birthday,
			   p.createdat as created_at,
			   p.updatedat as updated_at
		from patient p
		where id = $1;
`

	queryFetchPrescriptions = `
		select ps.id as id,
			   ps.stampID as stamp_id,
			   ps.typeID as type_id,
			   ps.statusID as status_id,
			   ps.medicinalProductID as medicinal_product_id,
			   mp.name as medicinal_product_name,
			   ps.medicinalProductQuantity as medicinal_product_quantity,
			   ps.doctorID as doctor_id,
			   ds.login as doctor_name,
			   ps.patientID as patient_id,
			   p.name as patient_name,
			   ps.pharmacistID as pharmacist_id,
			   phs.login as pharmacist_name,
			   ps.createdAt as created_at,
			   ps.updatedAt as updated_at,
			   ps.expiredAt as expired_at
		
		from prescriptions ps
			join medicinal_products mp
				on ps.medicinalProductID = mp.id
			join users ds
				on ps.doctorID = ds.id and ds.role_id = 1
			join patient p
				on ps.patientID = p.id
			join users phs
				 on ps.doctorID = phs.id and phs.role_id = 2
		order by id
		limit ($1 + 1) offset ($1 * ($2 - 1));
`

	queryGetPrescription = `
		select ps.id as id,
			   ps.stampID as stamp_id,
			   ps.typeID as type_id,
			   ps.statusID as status_id,
			   ps.medicinalProductID as medicinal_product_id,
			   mp.name as medicinal_product_name,
			   ps.medicinalProductQuantity as medicinal_product_quantity,
			   ps.doctorID as doctor_id,
			   ds.login as doctor_name,
			   ps.patientID as patient_id,
			   p.name as patient_name,
			   ps.pharmacistID as pharmacist_id,
			   phs.login as pharmacist_name,
			   ps.createdAt as created_at,
			   ps.updatedAt as updated_at,
			   ps.expiredAt as expired_at
		
		from prescriptions ps
			join medicinal_products mp
				on ps.medicinalProductID = mp.id
			join users ds
				on ps.doctorID = ds.id and ds.role_id = 1
			join patient p
				on ps.patientID = p.id
			join users phs
				 on ps.doctorID = phs.id and phs.role_id = 2
		where id = $1;
`
)
