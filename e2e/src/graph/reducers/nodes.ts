import _ from 'lodash';
import { getType } from 'typesafe-actions';

import { RootAction } from '../../rootAction';
import * as GraphActions from '../actions';
import { GraphNode } from '../types';


export interface NodeState {
  isLoading: boolean;
  nodes: GraphNode[];
  error: string;
};

export const InitialNodeState: NodeState = {
  isLoading: false,
  nodes: [],
  error: null,
};


export function NodesReducer(state: NodeState = InitialNodeState, action: RootAction) {
  switch(action.type) {

    /*              */
    /* CREATE NODE  */
    /*              */
    case getType(GraphActions['requestAddNode']):
      return {
        ...state,
        isLoading: true
      }
    case getType(GraphActions['successAddNode']):
      // clone array
      let cloned: GraphNode[] = [];
      state.nodes.forEach(node => { cloned.push(node) });
      cloned.push(action['payload']);
      return {
        ...state,
        isLoading: false,
        nodes: cloned
      }
    case getType(GraphActions['failureAddNode']):
      return {
        ...state,
        isLoading: false,
        error: action['payload']
      }

    case getType(GraphActions['requestRemoveNode']):
      return {
        ...state,
        isLoading: true,
      }
    case getType(GraphActions['successRemoveNode']):
      let aIndex: number | undefined = _.findIndex(state.nodes, action['payload'])
      return {
        ...state,
        isLoading: false,
        nodes: [...state.nodes.slice(0, aIndex),
                ...state.nodes.slice(aIndex+1)]
      }
    case getType(GraphActions['failureRemoveNode']):
      return {
        ...state,
        isLoading: false,
        error: action['payload'],
      }

    case getType(GraphActions['requestUpdateNode']):
      return {
        ...state,
        isLoading: true,
      }
    case getType(GraphActions['successUpdateNode']):
      let copy: GraphNode[] = [...state.nodes];
      let index: number = _.findIndex(copy, { img: action['payload']['prevImg'] });
      let newNode: GraphNode = { id: action['payload']['id'], img: action['payload']['img'], type: action['payload']['type'] }
      copy[index] = newNode
      return {
        ...state,
        isLoading: false,
        nodes: copy,
      }
    case getType(GraphActions['failureUpdateNode']):
      return {
        ...state,
        isLoading: false,
        error: action['payload'],
      }
    default: return state;
  }
};
