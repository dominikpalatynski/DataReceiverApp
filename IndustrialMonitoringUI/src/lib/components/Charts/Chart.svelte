<script lang="ts">
    import { onMount, onDestroy } from 'svelte';
    import { Chart, LineController, LineElement, PointElement, LinearScale, Title, CategoryScale } from 'chart.js/auto';

    Chart.register(LineController, LineElement, PointElement, LinearScale, Title, CategoryScale);

    type InfluxData = {
        time: string;
        value: number;
    };

    let chartData: InfluxData[] = [];
    let socket: WebSocket;

    let chart: Chart;

    let canvas: HTMLCanvasElement;

    function connectWebSocket() {
        socket = new WebSocket('ws://localhost:5050/ws');

        socket.onopen = () => {
            const config = {
                bucket: "org-bucket",
                variable_name: "temperature",
                measurement: "ESP1"
            };
            socket.send(JSON.stringify(config));
            console.log("Połączono z WebSocketem");
        };

        socket.onmessage = (event: MessageEvent) => {
            const parsedData = JSON.parse(event.data);

            console.log("Dane z WebSocket:", parsedData);
            // chartData = [...chartData, ...parsedData];

        };

        socket.onerror = (error) => {
            console.error("Błąd WebSocket:", error);
        };

        socket.onclose = () => {
            console.log("WebSocket zamknięty, ponowne łączenie za 5 sekund");
            setTimeout(connectWebSocket, 5000); // Ponowne połączenie po 5 sekundach
        };
    }

    function updateChart() {
        if (chart) {
            const labels = chartData.map((data) => data.time);
            const values = chartData.map((data) => data.value);

            chart.data.labels = labels;
            chart.data.datasets[0].data = values;
            chart.update();
        }
    }

    function initChart() {
        if (canvas) {
            chart = new Chart(canvas, {
                type: 'line',
                data: {
                    labels: [],
                    datasets: [
                        {
                            label: 'Influx Data',
                            data: [],
                            borderColor: 'rgba(75, 192, 192, 1)',
                            backgroundColor: 'rgba(75, 192, 192, 0.2)',
                            borderWidth: 1,
                            fill: true,
                        }
                    ]
                },
                options: {
                    responsive: true,
                    plugins: {
                        legend: {
                            display: true,
                            position: 'top'
                        }
                    },
                    scales: {
                        x: {
                            type: 'category',
                            title: {
                                display: true,
                                text: 'Time'
                            }
                        },
                        y: {
                            type: 'linear',
                            title: {
                                display: true,
                                text: 'Value'
                            }
                        }
                    }
                }
            });
        }
    }

    onMount(() => {
        const ctx = document.getElementById('chart') as HTMLCanvasElement
        chart = new Chart(ctx, {
            //Type of the chart
            type: 'line', 
            data: {
                //labels on x-axis
                labels: [], 
                datasets: [{
                    //The label for the dataset which appears in the legend and tooltips.
                    label: 'Price',
                    //data for the line
                    data: [],
                    //styling of the chart
                    backgroundColor: [
                        'rgba(255, 99, 132, 0.2)',
                    ],
                    borderColor: [
                        'rgba(255, 99, 132, 1)',
                    ],
                    borderWidth: 1
                }]
            },
            options: {
                scales: {
                    //make sure Y-axis starts at 0
                    y: {
                        beginAtZero: true
                    }
                },
            }
        });

        connectWebSocket();

        setTimeout(() => {
            const labels = chartData.map((data) => data.time);
            const values = chartData.map((data) => data.value);
            chart.update();
         }, 4000);

    });

    onDestroy(() => {
        if (socket) {
            socket.close();
        }
        if (chart) {
            chart.destroy();
        }
    });
</script>

<div>
    <canvas id="chart"></canvas>
</div>
<!-- <ul>
    {#each chartData as data}
        <li>{data.time}: {data.value}</li>
    {/each}
</ul> -->
