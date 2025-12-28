let chart = null;

function hitung() {
    const n = parseInt(document.getElementById("n").value);
    const nilaiInput = document.getElementById("nilai").value;

    const nilai = nilaiInput.split(",").map(Number);

    if (nilai.length < n) {
        alert("Jumlah nilai tidak sesuai dengan n!");
        return;
    }

    fetch("http://localhost:8080/hitung", {
        method: "POST",
        headers: {
            "Content-Type": "application/json"
        },
        body: JSON.stringify({
            n: n,
            nilai: nilai
        })
    })
        .then(res => res.json())
        .then(data => {
            document.getElementById("iteratif").innerText =
                `Rata-rata Iteratif : ${data.rata_iteratif.toFixed(2)} 
             (Waktu: ${data.waktu_iteratif} ns)`;

            document.getElementById("rekursif").innerText =
                `Rata-rata Rekursif : ${data.rata_rekursif.toFixed(2)} 
             (Waktu: ${data.waktu_rekursif} ns)`;

            tampilGrafik(n);
        });
}

function tampilGrafik(n) {
    const ctx = document.getElementById("grafik");

    let dataN = [];
    let iteratif = [];
    let rekursif = [];

    for (let i = 1; i <= n; i++) {
        dataN.push(i);
        iteratif.push(i);   // O(n)
        rekursif.push(i);   // O(n) + stack
    }

    if (chart) chart.destroy();

    chart = new Chart(ctx, {
        type: 'line',
        data: {
            labels: dataN,
            datasets: [
                {
                    label: 'Iteratif O(n)',
                    data: iteratif,
                    borderWidth: 2
                },
                {
                    label: 'Rekursif O(n)',
                    data: rekursif,
                    borderWidth: 2
                }
            ]
        },
        options: {
            plugins: {
                legend: {
                    position: 'top'
                }
            }
        }
    });
}