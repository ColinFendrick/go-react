import http from '../http-common';

const get = () => http.get('/task');

const service = { get };

export default service;
