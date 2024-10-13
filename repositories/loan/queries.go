package loan

const QUERY_GET_OUTSTANDING_BY_LOAN_ID = `
	SELECT COALESCE(SUM(ls.amount), 0) as outstanding_amount  FROM loan_schedules AS ls
	LEFT JOIN payments AS p ON p.loan_schedule_id = ls.id
	WHERE
		ls.loan_id = ? AND
		(p.status = 'pending' OR p.id IS NULL);
`

const QUERY_IS_DELINQUENT_BY_LOAN_ID = `
	SELECT 1 FROM (
		SELECT COUNT(1) AS pending_payments FROM loan_schedules AS ls
			LEFT JOIN payments AS p ON p.loan_schedule_id = ls.id
			WHERE
				ls.loan_id = ? AND
				ls.due_date <= CURRENT_TIMESTAMP AND
				(p.status = 'pending' OR p.id IS NULL)
	) as t
	WHERE t.pending_payments >= 2;
`
