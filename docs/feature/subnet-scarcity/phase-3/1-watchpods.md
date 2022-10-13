CNS current IPAM solution is reactive: it waits for the CNI to request (or release) an IP address for a Pod, and attempts to honor that request out of the current IP Pool. Because of this, requests are necessarily serial, and this creates a scaling bottleneck. 
When a Pod is created, the CNI will call with a request to assign an IP. If CNS is out of IPs and cannot honor that request, the CNI will return an error to the CRI, which will follow up by tearing down that Pod sandbox and starting over. Because of this stateless retrying, CNS can only reliable understand that it needs _at least one more_ IP, because it is impossible to tell if subsequent requests are retries for the same Pod, or many different Pods. If _many_ Pods have been scheduled, CNS will still only request a single additional batch of IPs, and assign those IPs one at a time until it runs out, then request a single additional batch of IPs...

A more predictive method of IP Pool scaling will be added to CNS: CNS will watch Pods for its Node, and will request/release IPs immediately based on the number of Pods scheduled. The Batching behavior will be unchanged, and CNS will continue to request IPs in Batches $B$ based on the local IP usage.