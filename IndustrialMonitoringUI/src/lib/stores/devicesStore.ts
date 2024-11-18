import type { Device } from '$lib/models/Device.js';
import { writable } from 'svelte/store';

const devicesStore = writable<Device[]>([]);

export default devicesStore;