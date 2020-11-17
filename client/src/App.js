import { useEffect } from 'react';
import { Switch, Route } from 'react-router-dom';

import AppService from './services/AppService';
import { Home } from './components';

const App = () => {
	useEffect(() => {
		(async () => {
			try {
				const res = await AppService.get();
				console.log(res);
			} catch (e) {
				console.log(e);
			}

		})();
	});
	return (
		<div>
      Basic Setup:
			<Switch>
				<Route path='/' exact component={Home} />
			</Switch>
		</div>

	);
};

export default App;
