import matplotlib.pyplot as plt

if __name__ == "__main__":
    import sys
    args = sys.argv[1:]
    if len(args) != 2:
        print("Usage: python visualization.py <log_filename> <output_image>")
        sys.exit(1)

    log_filename = args[0]
    output_image = args[1]

    latencies = []
    with open(log_filename, 'r') as f:
        for line in f:
            parts = line.strip().split(',')
            if len(parts) == 4:
                try:
                    latency = float(parts[1]) - float(parts[0])
                    latencies.append(latency)
                except ValueError:
                    continue

    if not latencies:
        print("No valid latency data found.")
        sys.exit(1)

    x_values = list(range(len(latencies)))
    
    plt.figure(figsize=(10, 5))
    plt.plot(x_values, latencies, marker='.', linestyle='None', markersize=2)
    plt.xlabel('N-th Request')
    plt.ylabel('Latency (ns)')
    plt.grid(True)

    plt.savefig(output_image, dpi=300, bbox_inches='tight')
