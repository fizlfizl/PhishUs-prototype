










AUI.add(
	'portal-available-languages',
	function(A) {
		var available = {};

		var direction = {};

		

			available['ar_SA'] = 'Arabic (Saudi Arabia)';
			direction['ar_SA'] = 'rtl';

		

			available['en_US'] = 'English (United States)';
			direction['en_US'] = 'ltr';

		

		Liferay.Language.available = available;
		Liferay.Language.direction = direction;
	},
	'',
	{
		requires: ['liferay-language']
	}
);