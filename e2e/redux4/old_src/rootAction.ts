import { ActionType } from 'typesafe-actions';

import * as DogActions from './dog/actions';
import * as GraphActions from './graph/actions';
import * as BookActions from './book/actions';
import * as TimeoutRetryActions from './timeoutRetryTest/actions';
import * as ProtocActions from 'protos/BasicState/actions_pb';
import * as ConfigActions from './config/actions'

const actions = {
  ...DogActions,
  ...GraphActions,
  ...BookActions,
  ...ProtocActions,
  ...ConfigActions,
  ...TimeoutRetryActions,
}

export type RootAction = ActionType<typeof actions>;

