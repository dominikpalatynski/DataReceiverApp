import type { Device } from '../models/Device.js';
import type { Slot } from '../models/Slot.js';
import type { Sensor } from '../models/Sensor.js';


export const fetchDevicesByOrgId = async(id: string): Promise<Device[]> => {
    try {
        const response = await fetch(`http://localhost:5000/devices/${id}`,{
            method: 'GET',
            headers: {
                'Content-Type': 'application'
            },
            credentials: 'include',
        }
            
        );
        if (!response.ok) throw new Error('Failed to fetch device');
      
        return await response.json();
    } catch (error) {
        console.error(error);
        throw error;
    }

}

export const fetchDeviceById = async(id: string): Promise<Device> => {
    try {
        const response = await fetch(`http://localhost:5000/org/devices/${id}`, {
            method: 'GET',
            headers: {
                'Content-Type': 'application/json'
            }
        });
        if (!response.ok) throw new Error('Failed to fetch device');
      
        return await response.json();
    } catch (error) {
        console.error(error);
        throw error;
    }
}

export const addDeviceInfo = async(device: {name: string; interval: number; org_id: number}): Promise<Device> => {
    try {
        const response = await fetch('http://localhost:5000/devices', {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json'
            },
            body: JSON.stringify(device)
        });
        if (!response.ok) throw new Error('Failed to add device');
      
        return await response.json();
    } catch (error) {
        console.error(error);
        throw error;
    }

}

export const fetchSlotsByDeviceId = async(id: string): Promise<Slot[]> => {
    try {
        const response = await fetch(`http://localhost:5000/deviceData/slots/${id}`,{
            method: 'GET',
            headers: {
                'Content-Type': 'application/json'
            },
            credentials: 'include',
        }
            
        );
        if (!response.ok) throw new Error('Failed to fetch slots');
      
        return await response.json();
    } catch (error) {
        console.error(error);
        throw error;
    }

}

export const addSensor = async(sensor: {device_id: number; name: string; variable_name: string; slot: number}): Promise<Sensor> => {
    try {
        const response = await fetch('http://localhost:5000/device/sensor/add', {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json'
            },
            credentials: 'include',
            body: JSON.stringify(sensor)
        });
        if (!response.ok) throw new Error('Failed to add sensor');
      
        return await response.json();
    } catch (error) {
        console.error(error);
        throw error;
    }

}