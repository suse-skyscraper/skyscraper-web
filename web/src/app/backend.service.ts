import { Injectable } from '@angular/core';
import { HttpClient } from '@angular/common/http';
import { Observable } from 'rxjs';
import config from './app.config';

@Injectable({
  providedIn: 'root',
})
export class BackendService {
  private readonly host: string;

  constructor(private http: HttpClient) {
    this.host = config.backend.host;
  }

  getProfile(): Observable<UserResponse> {
    const url = new URL('/api/v1/profile', this.host);
    return this.http.get<UserResponse>(url.href);
  }

  getCloudTenants(): Observable<CloudTenantsResponse> {
    const url = new URL('/api/v1/cloud_tenants', this.host);
    return this.http.get<CloudTenantsResponse>(url.href);
  }

  getCloudTenantTags(
    cloud: string,
    tenantId: string,
  ): Observable<CloudTenantTags> {
    const url = new URL(
      `/api/v1/cloud_tenants/cloud/${cloud}/tenant/${tenantId}/tags`,
      this.host,
    );
    return this.http.get<CloudTenantTags>(url.href);
  }

  getCloudAccount(
    cloud: string,
    tenantId: string,
    accountId: string,
  ): Observable<CloudAccountResponse> {
    const url = new URL(
      `/api/v1/cloud_tenants/cloud/${cloud}/tenant/${tenantId}/accounts/${accountId}`,
      this.host,
    );
    return this.http.get<CloudAccountResponse>(url.href);
  }

  updateCloudAccount(
    cloud: string,
    tenantId: string,
    accountId: string,
    update: UpdateCloudAccountRequest,
  ): Observable<CloudAccountResponse> {
    const url = new URL(
      `/api/v1/cloud_tenants/cloud/${cloud}/tenant/${tenantId}/accounts/${accountId}`,
      this.host,
    );
    return this.http.put<CloudAccountResponse>(url.href, update);
  }

  getCloudAccounts(
    cloud: string,
    tenantId: string,
    filter?: Map<string, string>,
  ): Observable<CloudAccountsResponse> {
    const url = new URL(
      `/api/v1/cloud_tenants/cloud/${cloud}/tenant/${tenantId}/accounts`,
      this.host,
    );
    if (filter !== undefined) {
      filter.forEach((value, key) => {
        url.searchParams.append(key, value);
      });
    }
    return this.http.get<CloudAccountsResponse>(url.href);
  }
}

export interface UserAttributes {
  username: string;
}

export interface UserItem {
  id: string;
  type: string;
  attributes: UserAttributes;
}

export interface UserResponse {
  data: UserItem;
}

export interface CloudTenantItem {
  cloud_provider: string;
  tenant_id: string;
  name: string;
  active: boolean;
  created_at: string;
  updated_at: string;
}

export interface CloudTenantsResponse {
  data: CloudTenantItem[];
}

export interface CloudAccountResponse {
  data: CloudAccountItem | null;
}

export interface CloudAccountsResponse {
  data: CloudAccountItem[];
}

export interface CloudAccountItem {
  id: string;
  type: string;
  attributes: CloudAccountAttributes;
}

export interface CloudAccountAttributes {
  cloud_provider: string;
  tenant_id: string;
  account_id: string;
  name: string;
  active: boolean;
  tags_desired: { [key: string]: string };
  tags_current: { [key: string]: string };
  tags_drift_detected: boolean;
  created_at: string;
  updated_at: string;
}

export interface UpdateCloudAccountRequest {
  data: UpdateCloudAccountRequestData;
}

export interface UpdateCloudAccountRequestData {
  tags_desired: { [key: string]: string };
}

export interface CloudTenantTags {
  tags: string[];
}
