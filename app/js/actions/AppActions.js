/*
 * Actions change things in your application
 * Since this boilerplate uses a uni-directional data flow, specifically redux,
 * we have these actions which are the only way your application interacts with
 * your appliction state. This guarantees that your state is up to date and nobody
 * messes it up weirdly somewhere.
 *
 * To add a new Action:
 * 1) Import your constant
 * 2) Add a function like this:
 *    export function yourAction(var) {
 *        return { type: YOUR_ACTION_CONSTANT, var: var }
 *    }
 * 3) (optional) Add an async function like this:
 *    export function asyncYourAction(var) {
 *        return (dispatch) => {
 *             // Do async stuff here
 *             return dispatch(yourAction(var));
 *        };
 *    }
 *
 *    If you add an async function, remove the export from the function
 *    created in the second step
 */

// Disable the no-use-before-define eslint rule for this file
// It makes more sense to have the asnyc actions before the non-async ones
/* eslint-disable no-use-before-define */

import * as AppConstants from '../constants/AppConstants';

//export function asyncChangeProjectName(name) {
//  return (dispatch) => {
//    // You can do async stuff here!
//    // API fetching, Animations,...
//    // For more information as to how and why you would do this, check https://github.com/gaearon/redux-thunk
//    return dispatch(changeProjectName(name));
//  };
//}
//
//export function asyncChangeOwnerName(name) {
//  return (dispatch) => {
//    // You can do async stuff here!
//    // API fetching, Animations,...
//    // For more information as to how and why you would do this, check https://github.com/gaearon/redux-thunk
//    return dispatch(changeOwnerName(name));
//  };
//}
//
//export function changeProjectName(name) {
//  return {type: AppConstants.CHANGE_PROJECT_NAME, name};
//}
//
//export function changeOwnerName(name) {
//  return {type: AppConstants.CHANGE_OWNER_NAME, name};
//}

export function deleteStruct(struct) {
  return {
    type: AppConstants.DELETE_STRUCT,
    struct,
  };
}

export function changeStructName(struct) {
  return {
    type: AppConstants.CHANGE_STRUCT_NAME,
    struct,
  };
}

export function changeStructFieldName(struct) {
  return {
    type: AppConstants.CHANGE_STRUCT_FIELD_NAME,
    struct,
  };
}

export function changeStructFieldType(struct) {
  return {
    type: AppConstants.CHANGE_STRUCT_FIELD_TYPE,
    struct,
  };
}

export function addStructField(struct) {
  return {
    type: AppConstants.ADD_STRUCT_FIELD,
    struct,
  };
}

export function removeStructField(struct) {
  return {
    type: AppConstants.REMOVE_STRUCT_FIELD,
    struct,
  };
}

export function setPackageData(packageData) {
  return {
    type: AppConstants.SET_PACKAGE_DATA,
    packageData,
  };
}

export function onAddStruct(file) {
  return {
    type: AppConstants.ADD_STRUCT,
    file
  }
}
