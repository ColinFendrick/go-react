import { Switch, Route } from 'react-router-dom';

import { Home } from './comonents';

const App = () => {
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
