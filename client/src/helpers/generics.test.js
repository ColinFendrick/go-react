import {
	pipe
} from './generics';

describe('Testing pipe fn', () => {
	test('Works if the functions are good', () => {
		const f = n => n * 2;
		const g = n => n + 1;
		const h = n => n - 4;

		expect(pipe(h, g, f)(9)).toEqual(12);
		expect(pipe(f, g, h)(9)).toEqual(15);
	});

	test('Errors otherwise', () => {
		const f = n => n * 2;
		const g = 6;

		expect(() => pipe(f, g)(20)).toThrow(TypeError);
	});
});
