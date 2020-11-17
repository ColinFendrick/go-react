import '@testing-library/jest-dom';
import { render } from '@testing-library/react';
import { BrowserRouter } from 'react-router-dom';

import ContextContainer from './containers/ContextContainer';
import { pipe } from './helpers/generics';

const withWrapper = (Wrapper, wrapperProps = {}) => children =>
	<Wrapper {...wrapperProps}>{children}</Wrapper>;

export const withRouter = withWrapper(BrowserRouter);
export const withContext = withWrapper(ContextContainer);

export const renderWith = (...wrappers) => children => render(
	pipe(...wrappers)(children)
);

export const setup = (...builtins) => (...fns) =>
	builtins.forEach((builtin, ix) =>
		builtin(() =>
			fns[0].length ? fns[ix].forEach(f => f()) : fns.forEach(f => f())
		)
	);
