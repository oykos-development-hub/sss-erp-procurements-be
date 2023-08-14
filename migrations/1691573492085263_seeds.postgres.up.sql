--planovi, po jedan plan za svaku godinu, postoji i jedan postbudzetski
INSERT INTO plans (
    year, title, active, serial_number, date_of_publishing, date_of_closing, pre_budget_id, file_id, created_at,updated_at)
VALUES
    ('2022', 'Plan 2022', false, 'PN001', '2022-01-15', '2022-02-15', null, 101, NOW(), NOW()),
    ('2023', 'Plan 2023', false, 'PN002', '2023-12-01', '2024-01-31', null, 102, NOW(), NOW()),
    ('2024', 'Plan 2024', true, 'PN003', '2024-01-01', '2024-02-28', null, 103, NOW(), NOW()),
    ('2023', 'Plan 2023', true, 'PN004', '2023-12-01', '2024-01-31', 2, 102, NOW(), NOW());


--items za planove, prva 4 su za prebudzetske a poslednja 2 za postbudzetski plan
INSERT INTO items (
	budget_indent_id, plan_id, is_open_procurement, title, article_type, status, serial_number, date_of_publishing, date_of_awarding, file_id, created_at, updated_at
)
VALUES
    (1, 1, true, 'Namjestaj', 'Roba', 'U toku', 'SN001', '2023-03-01', '2023-03-15', 101, NOW(), NOW()),
    (1, 1, true, 'Struja', 'Usluga', 'U toku', 'SN002', '2023-04-01', null, 102, NOW(), NOW()),
    (2, 2, true, 'Namjestaj', 'Roba', 'Aktivan', 'SN005', '2023-04-01', null, 102, NOW(), NOW()),
    (2, 2, false, 'Telefonija', 'Usluga', 'Aktivan', 'SN006', '2023-05-01', '2023-05-10', 103, NOW(), NOW()),
    (4, 4, true, 'Namjestaj', 'Roba', 'Aktivan', 'SN007', '2023-03-01', '2023-03-15', 101, NOW(), NOW()),
    (4, 4, true, 'Struja', 'Usluga', 'Aktivan', 'SN008', '2023-04-01', null, 102, NOW(), NOW());

--artikli, po istom principu, i za svaki namjestaj imamo 2 artikla 
INSERT INTO articles (
    budget_indent_id, item_id, title,description, net_price, vat_percentage, created_at, updated_at
)
VALUES
    (1, 1, 'Kauc', 'Namjestaj kauc', '100.00', '21', NOW(), NOW()),
    (1, 1, 'Stolica', 'Namjestaj stolica', '50.00', '21', NOW(), NOW()),
    (1, 2, 'Racun', 'Racun za struju', '200.00', '21', NOW(), NOW()),
    (2, 3, 'Sto', 'Namjestaj sto', '100.00', '21', NOW(), NOW()),
    (2, 3, 'Polica', 'Namjestaj polica', '100.00', '21', NOW(), NOW()),
    (2, 4, 'Racun', 'Racun za telefon', '200.00', '21', NOW(), NOW()),
    (3, 5, 'Kauc', 'Namjestaj kauc', '100.00', '21', NOW(), NOW()),
    (3, 5, 'Stolica', 'Namjestaj stolica', '50.00', '21', NOW(), NOW()),
    (3, 6, 'Racun', 'Racun za struju', '200.00', '21', NOW(), NOW());
	

--ugovori sa dobavljacima
INSERT INTO contracts (
	public_procurement_id, supplier_id, serial_number, date_of_signing, date_of_expiry, net_value, gross_value, created_at, updated_at, file_id
)
VALUES
    (1, 1, 'CON001', '2023-03-01', '2024-03-01', '10000.00', '12000.00', NOW(), NOW(), 101),
    (2, 2, 'CON002', '2023-04-01', '2024-04-01', '5000.00', '6000.00', NOW(), NOW(), 102),
    (3, 1, 'CON003', '2023-05-01', null, '15000.00', '18000.00', NOW(), NOW(), 103),
	(4, 3, 'CON001', '2023-03-01', '2024-03-01', '10000.00', '12000.00', NOW(), NOW(), 101),
    (5, 1, 'CON002', '2023-04-01', '2024-04-01', '5000.00', '6000.00', NOW(), NOW(), 102),
    (6, 2, 'CON003', '2023-05-01', null, '15000.00', '18000.00', NOW(), NOW(), 103);


--zahtjevi organizacionih jedinica, ima prihvacenih ima i odbijenih
INSERT INTO organization_unit_articles (
	article_id, organization_unit_id, amount, status, is_rejected,rejected_description, created_at,updated_at
)
VALUES
    (1, 1, 1, '', false, null, NOW(), NOW()),
    (1, 2, 2, 'rejected', true, 'Prevelika investicija', NOW(), NOW()),
    (2, 1, 1, '', false, null, NOW(), NOW()),
	(3, 3, 2, 'rejected', true, 'Prevelika investicija', NOW(), NOW()),
    (3, 3, 1, '', false, null, NOW(), NOW());
	
--limiti za pojedine organizacione jedinice
INSERT INTO organization_unit_plan_limits (
	organization_unit_id, item_id,limit_value,created_at,  updated_at
)
VALUES
    (1, 1, 500, NOW(), NOW()),
    (1, 2, 500, NOW(), NOW()),
    (2, 2, 200, NOW(), NOW()),
    (3, 4, 200, NOW(), NOW()),
    (3, 5, 200, NOW(), NOW());

--spojeni ugovori i artikli
INSERT INTO contract_articles (
	article_id, contract_id, amount, net_value, gross_value, created_at,updated_at
)
VALUES
    (1, 1, 1, '1000.00', '1200.00', NOW(), NOW()),
    (3, 2, 3, '1000.00', '1200.00', NOW(), NOW()),
    (4, 3, 2, '200.00', '260.00', NOW(), NOW()),
	(6, 4, 3, '1000.00', '1200.00', NOW(), NOW()),
    (7, 5, 2, '200.00', '260.00', NOW(), NOW());
