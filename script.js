let chart;

function proses() {
    const input = document.getElementById("nilai").value.trim();

    if (input === "") {
        alert("Nilai tidak boleh kosong!");
        return;
    }

    const nilai = input
        .split(",")
        .map(n => parseInt(n.trim()))
        .filter(n => !isNaN(n)); // ⬅️ PENTING

    if (nilai.length === 0) {
        alert("Format salah! Contoh: 80,75,90");
        return;
    }

    fetch("/proses", {
        method: "POST",
        headers: { "Content-Type": "application/json" },
        body: JSON.stringify({ nilai: nilai })
    })
        .then(res => {
            if (!res.ok) throw new Error("Server error");
            return res.json();
        })
        .then(data => {

            document.getElementById("rataIter").innerText =
                data.rata_iteratif.toFixed(2);

            document.getElementById("rataRek").innerText =
                data.rata_rekursif.toFixed(2);

            document.getElementById("timeIter").innerText =
                data.time_iteratif;

            document.getElementById("timeRek").innerText =
                data.time_rekursif;

            const ctx = document.getElementById("chart");

            if (chart) chart.destroy();

            chart = new Chart(ctx, {
                type: "bar",
                data: {
                    labels: ["Iteratif", "Rekursif"],
                    datasets: [{
                        label: "Running Time (ns)",
                        data: [data.time_iteratif, data.time_rekursif]
                    }]
                }
            });
        })
        .catch(err => {
            alert("Gagal memproses data!");
            console.error(err);
        });
}
