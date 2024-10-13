package model

const queryUpdateLoanStatusToCompletedHook = `
	UPDATE loans
	SET status = 'completed'
	WHERE
		loans.id = (SELECT loan_id FROM loan_schedules WHERE id = ?)
	AND NOT EXISTS (
		SELECT 1
		FROM payments p
		WHERE p.loan_schedule_id IN (SELECT id FROM loan_schedules WHERE loan_id = loans.id)
		AND p.status != 'paid'
	);
`
