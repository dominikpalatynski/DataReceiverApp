export interface Org {
    id: number;
    name: string;
    bucket: string;
  }

export interface OrganizationName {
    id: number;
    name: string;
  }
  
export interface UserOrganization {
    organization: OrganizationName;
    role: string;
  }