import { createAction } from 'typesafe-actions';
import { GraphNode, Label, UpdateLabel, UpdateGraphNode } from './types';

// Add Node
export const requestAddNode = createAction('GRAPH_REQUEST_ADD_NODE');
export const successAddNode = createAction('GRAPH_ADD_NODE', (resolve) => {
  return (node: GraphNode) => resolve(node);
});
export const failureAddNode = createAction('GRAPH_FAILURE_ADD_NODE', (resolve) => {
  return (error: string) => resolve(error);
});

export const requestRemoveNode = createAction('GRAPH_REQUEST_REMOVE_NODE', (resolve) => {
  return (node: GraphNode) => resolve(node);
});
export const successRemoveNode = createAction('GRAPH_REMOVE_NODE', (resolve) => {
  return (node: GraphNode) => resolve(node);
});
export const failureRemoveNode = createAction('GRAPH_FAILURE_REMOVE_NODE', (resolve) => {
  return (error: string) => resolve(error);
});


export const requestUpdateNode = createAction('GRAPH_REQUEST_UPDATE_NODE', (resolve) => {
  return (node: GraphNode) => resolve(node);
});
export const successUpdateNode = createAction('GRAPH_UPDATE_NODE', (resolve) => {
  return (node: UpdateGraphNode) => resolve(node);
});
export const failureUpdateNode = createAction('GRAPH_FAILURE_UPDATE_NODE', (resolve) => {
  return (error: string) => resolve(error);
});

// Add Label (map of label name to nodeids)
export const createLabelRequest = createAction('GRAPH_CREATE_LABEL_REQUEST')
export const createLabelSuccess = createAction('GRAPH_CREATE_LABEL_SUCCESS', (resolve) => {
  return (label: Label) => resolve(label);
});
export const createLabelFailure = createAction('GRAPH_CREATE_LABEL_FAILURE', (resolve) => {
  return (error: string) => resolve(error);
});

export const updateLabelRequest = createAction('GRAPH_UPDATE_LABEL_REQUEST', (resolve) => {
  return (label: Label) => resolve(label);
});
export const updateLabelSuccess = createAction('GRAPH_UPDATE_LABEL_SUCCESS', (resolve) => {
  return (label: UpdateLabel) => resolve(label);
});
export const updateLabelFailure = createAction('GRAPH_UPDATE_LABEL_FAILURE', (resolve) => {
  return (error: string) => resolve(error);
});

export const deleteLabelRequest = createAction('GRAPH_DELETE_LABEL_REQUEST', (resolve) => {
  return (label: Label) => resolve(label);
});
export const deleteLabelSuccess = createAction('GRAPH_DELETE_LABEL_SUCCESS', (resolve) => {
  return (label: Label) => resolve(label);
});
export const deleteLabelFailure = createAction('GRAPH_DELETE_LABEL_FAILURE', (resolve) => {
  return (error: string) => resolve(error);
});
