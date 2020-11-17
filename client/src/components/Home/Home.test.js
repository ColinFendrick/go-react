import { screen } from '@testing-library/react';

import { renderWith, setup } from '../../setupTests';

import Home from './Home';

describe('Testing <Home />', () => {
	setup(beforeEach)(
		() => renderWith()(<Home />)
	);

	test('Renders the Pickers component given no data', () => {
		expect(
			screen.getByText('React Go')
		).toBeInTheDocument();
	});
});
