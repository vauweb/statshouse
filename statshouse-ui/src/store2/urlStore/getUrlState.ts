// Copyright 2024 V Kontakte LLC
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.

import { arrToObj, getDefaultParams, type QueryParams, toTreeObj, urlDecode } from 'url2';
import { UrlStore } from './urlStore';
import { resetDefaultParams } from './resetDefaultParams';
import { loadDashboard } from './loadDashboard';
import { type Location } from 'history';

export async function getUrlState(
  prevParam: QueryParams,
  location: Location
): Promise<Pick<UrlStore, 'params' | 'saveParams'> & { reset: boolean; error?: Error }> {
  const urlSearchArray = [...new URLSearchParams(location.search || location.hash.slice(1))];
  const urlObject = arrToObj(urlSearchArray);
  const urlTree = toTreeObj(urlObject);
  const { params: saveParams, error } = await loadDashboard(prevParam, urlTree, getDefaultParams());
  const params = urlDecode(urlTree, saveParams);
  const resetParams = resetDefaultParams(params);
  return {
    params: resetParams ?? params,
    saveParams,
    reset: !!resetParams,
    error,
  };
}