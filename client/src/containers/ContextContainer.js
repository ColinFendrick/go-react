import { AppProvider } from '../context/AppContext';

const ContextContainer = ({ children }) => (
	<AppProvider>
		{ children }
	</AppProvider>
);

export default ContextContainer;
