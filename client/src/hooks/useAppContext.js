import { useContext } from 'react';

import { AppContext } from '../context/AppContext';

const useAppContext = () => {
	const [appState, setAppState] = useContext(AppContext);

	return {
		appState, setAppState
	};
};

export default useAppContext;
