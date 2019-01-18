import { createAction } from 'typesafe-actions';

export const dogRequest = createAction('DOG_REQUEST')

export const dogCancel = createAction('DOG_CANCEL')

export const dogSuccess = createAction('DOG_SUCCESS', (resolve) => {
  return (dog: number) => resolve(dog);
})

export const dogFailure = createAction('DOG_FAILURE', (resolve) => {
  return (error: string) => resolve(error);
})
