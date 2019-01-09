import { createAction } from 'typesafe-actions';

export const SetHost = createAction('SET_HOST', (resolve) => {
  return (host: string) => resolve(host);
});
