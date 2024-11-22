import type { Org, UserOrganization } from "$lib/models/Org.js";


export const createOrg = async(org: {name: string; bucket: string;}): Promise<Org> => {
    try {
        const response = await fetch('http://localhost:5000/auth/org/create', {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json'
            },
            credentials: 'include',
            body: JSON.stringify(org)
        });
        if (!response.ok) throw new Error('Failed to create org');
      
        return await response.json();
    } catch (error) {
        console.error(error);
        throw error;
    }

}

export const fetchOrgsConnectedToUser = async(): Promise<UserOrganization[]> => {
    try {
        const response = await fetch(`http://localhost:5000/auth/org/connected`,{
            method: 'GET',
            headers: {
                'Content-Type': 'application'
            },
            credentials: 'include',
        }
            
        );
        if (!response.ok) throw new Error('Failed to fetch Orgs connected to user');
      
        return await response.json();
    } catch (error) {
        console.error(error);
        throw error;
    }

}