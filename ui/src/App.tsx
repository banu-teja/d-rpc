import { useState, useEffect } from 'react'
import './App.css'
import { ethers } from 'ethers'

// API configuration
const API_ENDPOINT = import.meta.env.VITE_API_URL || 'http://localhost:8080';

interface Provider {
  address: string
  qosScore: string
  stake: string
}

interface DiscoveryResponse {
  providers: Provider[]
  recommendedProvider: string
}

function App() {
  // State variables
  const [providers, setProviders] = useState<Provider[]>([])
  const [recommendedProvider, setRecommendedProvider] = useState<string>('')
  const [loading, setLoading] = useState(true)
  const [error, setError] = useState<string | null>(null)
  const [selectedProvider, setSelectedProvider] = useState<string>('')
  const [rpcEndpoint, setRpcEndpoint] = useState<string>(API_ENDPOINT)
  const [blockNumber, setBlockNumber] = useState<number | null>(null)
  const [channelId, setChannelId] = useState<string>('')
  const [isChannelOpen, setIsChannelOpen] = useState(false)
  const [activeTab, setActiveTab] = useState('dashboard')
  const [channelBalance, setChannelBalance] = useState('0.1')
  const [gasPrice, setGasPrice] = useState<string | null>(null)
  const [networkInfo, setNetworkInfo] = useState<{ name: string, chainId: number } | null>(null)
  const [requestCount, setRequestCount] = useState(0)

  // Fetch providers on component mount
  useEffect(() => {
    fetchProviders()
    const interval = setInterval(fetchProviders, 30000) // Refresh every 30 seconds
    return () => clearInterval(interval)
  }, [])

  // Update block number and network info when provider is selected
  useEffect(() => {
    if (selectedProvider) {
      fetchBlockNumber()
      fetchGasPrice()
      fetchNetworkInfo()
      const interval = setInterval(() => {
        fetchBlockNumber()
        fetchGasPrice()
      }, 10000)
      return () => clearInterval(interval)
    }
  }, [selectedProvider])

  const fetchProviders = async () => {
    try {
      setLoading(true)
      setError(null)
      const response = await fetch(`${rpcEndpoint}/discovery`, {
        headers: {
          'Accept': 'application/json',
          'Content-Type': 'application/json'
        },
        mode: 'cors'
      })
      if (!response.ok) {
        throw new Error(`Failed to fetch providers: ${response.statusText}`)
      }
      const data: DiscoveryResponse = await response.json()
      setProviders(data.providers)
      setRecommendedProvider(data.recommendedProvider)

      // Auto-select recommended provider if available
      if (data.recommendedProvider && !selectedProvider) {
        setSelectedProvider(data.recommendedProvider)
      }
    } catch (err) {
      console.error("Fetch error:", err)
      setError(`Error fetching providers: ${err instanceof Error ? err.message : String(err)}`)
    } finally {
      setLoading(false)
    }
  }

  const fetchBlockNumber = async () => {
    try {
      const provider = new ethers.JsonRpcProvider(`${rpcEndpoint}`)
      const blockNumber = await provider.getBlockNumber()
      setBlockNumber(blockNumber)
      // Simulate request count increasing
      setRequestCount(prev => prev + 1)
    } catch (err) {
      console.error("Error fetching block number:", err)
    }
  }

  const fetchGasPrice = async () => {
    try {
      const provider = new ethers.JsonRpcProvider(`${rpcEndpoint}`)
      const gasPrice = await provider.getFeeData()
      setGasPrice(ethers.formatUnits(gasPrice.gasPrice || 0, 'gwei'))
    } catch (err) {
      console.error("Error fetching gas price:", err)
    }
  }

  const fetchNetworkInfo = async () => {
    try {
      const provider = new ethers.JsonRpcProvider(`${rpcEndpoint}`)
      const network = await provider.getNetwork()
      setNetworkInfo({
        name: network.name === 'unknown' ? 'Anvil Local' : network.name,
        chainId: Number(network.chainId)
      })
    } catch (err) {
      console.error("Error fetching network info:", err)
    }
  }

  const openPaymentChannel = async () => {
    try {
      // This would call the smart contract to open a payment channel
      // For demo purposes, we're just setting a fake channel ID
      setChannelId(`channel_${Date.now().toString(16)}`)
      setIsChannelOpen(true)
    } catch (err) {
      setError(`Error opening payment channel: ${err instanceof Error ? err.message : String(err)}`)
    }
  }

  const closePaymentChannel = async () => {
    try {
      // This would call the smart contract to close the payment channel
      // For demo purposes, we're just resetting the UI
      setChannelId('')
      setIsChannelOpen(false)
    } catch (err) {
      setError(`Error closing payment channel: ${err instanceof Error ? err.message : String(err)}`)
    }
  }

  const handleProviderSelect = (providerAddress: string) => {
    setSelectedProvider(providerAddress)
  }

  const simulateTransaction = async () => {
    try {
      // Simulate a transaction being processed
      setLoading(true)
      await new Promise(resolve => setTimeout(resolve, 1000))
      setRequestCount(prev => prev + 1)
      setChannelBalance(prev => (parseFloat(prev) - 0.01).toFixed(2))
      setLoading(false)
    } catch (err) {
      setError(`Error processing transaction: ${err instanceof Error ? err.message : String(err)}`)
      setLoading(false)
    }
  }

  return (
    <div className="app">
      <header>
        <h1>Decentralized RPC Network</h1>
        <p>A network of incentivized RPC providers</p>

        <nav className="main-nav">
          <ul>
            <li className={activeTab === 'dashboard' ? 'active' : ''}>
              <button onClick={() => setActiveTab('dashboard')}>Dashboard</button>
            </li>
            <li className={activeTab === 'providers' ? 'active' : ''}>
              <button onClick={() => setActiveTab('providers')}>Providers</button>
            </li>
            <li className={activeTab === 'payments' ? 'active' : ''}>
              <button onClick={() => setActiveTab('payments')}>Payments</button>
            </li>
            <li className={activeTab === 'transactions' ? 'active' : ''}>
              <button onClick={() => setActiveTab('transactions')}>Transactions</button>
            </li>
          </ul>
        </nav>
      </header>

      <main>
        {error && (
          <div className="error-message">
            {error}
            <button onClick={fetchProviders}>Retry</button>
          </div>
        )}

        {activeTab === 'dashboard' && (
          <>
            <section className="network-stats">
              <h2>Network Status</h2>
              <div className="stats-container">
                <div className="stat-card">
                  <h3>Active Providers</h3>
                  <p>{providers.length}</p>
                </div>
                {blockNumber !== null && (
                  <div className="stat-card">
                    <h3>Current Block</h3>
                    <p>{blockNumber}</p>
                  </div>
                )}
                {gasPrice !== null && (
                  <div className="stat-card">
                    <h3>Gas Price</h3>
                    <p>{gasPrice} Gwei</p>
                  </div>
                )}
                {networkInfo && (
                  <div className="stat-card">
                    <h3>Network</h3>
                    <p>{networkInfo.name} (Chain ID: {networkInfo.chainId})</p>
                  </div>
                )}
                <div className="stat-card">
                  <h3>Payment Channel</h3>
                  <p>{isChannelOpen ? 'Open' : 'Closed'}</p>
                  {isChannelOpen && (
                    <span className="channel-id">ID: {channelId}</span>
                  )}
                </div>
                <div className="stat-card">
                  <h3>Requests Processed</h3>
                  <p>{requestCount}</p>
                </div>
              </div>
            </section>

            <section className="actions-panel">
              <h2>Quick Actions</h2>
              <div className="actions-container">
                <button
                  className="action-button"
                  onClick={fetchBlockNumber}
                  disabled={!selectedProvider}
                >
                  Fetch Latest Block
                </button>
                <button
                  className="action-button"
                  onClick={simulateTransaction}
                  disabled={!isChannelOpen}
                >
                  Simulate Transaction
                </button>
                {!isChannelOpen ? (
                  <button
                    className="action-button primary"
                    onClick={openPaymentChannel}
                  >
                    Open Payment Channel
                  </button>
                ) : (
                  <button
                    className="action-button danger"
                    onClick={closePaymentChannel}
                  >
                    Close Payment Channel
                  </button>
                )}
              </div>
            </section>
          </>
        )}

        {activeTab === 'providers' && (
          <section className="provider-section">
            <h2>RPC Providers</h2>

            {loading ? (
              <div className="loading">Loading providers...</div>
            ) : (
              <>
                {recommendedProvider && (
                  <div className="recommended-provider">
                    <h3>Recommended Provider</h3>
                    <div
                      className={`provider-card ${selectedProvider === recommendedProvider ? 'selected' : ''}`}
                      onClick={() => handleProviderSelect(recommendedProvider)}
                    >
                      <div className="provider-address">{recommendedProvider}</div>
                      <div className="provider-metrics">
                        {providers.find(p => p.address === recommendedProvider)?.qosScore ? (
                          <>
                            <span>QoS Score: {providers.find(p => p.address === recommendedProvider)?.qosScore}</span>
                            <span>Stake: {
                              providers.find(p => p.address === recommendedProvider)?.stake
                                ? ethers.formatEther(providers.find(p => p.address === recommendedProvider)?.stake || '0')
                                : '0'
                            } ETH</span>
                          </>
                        ) : (
                          'Loading metrics...'
                        )}
                      </div>
                      {selectedProvider === recommendedProvider && (
                        <div className="selected-badge">Currently Selected</div>
                      )}
                    </div>
                  </div>
                )}

                <h3>All Providers</h3>
                <div className="providers-list">
                  {providers.length === 0 ? (
                    <p>No providers available</p>
                  ) : (
                    providers.map((provider) => (
                      <div
                        key={provider.address}
                        className={`provider-card ${selectedProvider === provider.address ? 'selected' : ''}`}
                        onClick={() => handleProviderSelect(provider.address)}
                      >
                        <div className="provider-address">{provider.address}</div>
                        <div className="provider-metrics">
                          <span>QoS Score: {provider.qosScore}</span>
                          <span>Stake: {ethers.formatEther(provider.stake)} ETH</span>
                          <span>Estimated Uptime: {Math.floor(Math.random() * 10) + 90}%</span>
                        </div>
                        {selectedProvider === provider.address && (
                          <div className="selected-badge">Currently Selected</div>
                        )}
                      </div>
                    ))
                  )}
                </div>
              </>
            )}
          </section>
        )}

        {activeTab === 'payments' && (
          <section className="payment-section">
            <h2>Payment Channel Management</h2>
            {!isChannelOpen ? (
              <div className="channel-creation">
                <h3>Open a New Payment Channel</h3>
                <div className="form-group">
                  <label>Provider Address</label>
                  <input type="text" value={selectedProvider} readOnly />
                </div>
                <div className="form-group">
                  <label>Initial Deposit (ETH)</label>
                  <input
                    type="number"
                    min="0.01"
                    step="0.01"
                    value={channelBalance}
                    onChange={(e) => setChannelBalance(e.target.value)}
                  />
                </div>
                <div className="form-group">
                  <label>Duration</label>
                  <select>
                    <option>1 Hour</option>
                    <option>1 Day</option>
                    <option selected>1 Week</option>
                    <option>1 Month</option>
                  </select>
                </div>
                <button className="open-channel-btn" onClick={openPaymentChannel}>
                  Open Payment Channel
                </button>
              </div>
            ) : (
              <div className="channel-details">
                <h3>Active Payment Channel</h3>
                <div className="channel-info">
                  <div className="info-row">
                    <span>Channel ID:</span>
                    <span>{channelId}</span>
                  </div>
                  <div className="info-row">
                    <span>Provider:</span>
                    <span>{selectedProvider}</span>
                  </div>
                  <div className="info-row">
                    <span>Status:</span>
                    <span className="channel-active">Active</span>
                  </div>
                  <div className="info-row">
                    <span>Balance:</span>
                    <span>{channelBalance} ETH</span>
                  </div>
                  <div className="info-row">
                    <span>Created:</span>
                    <span>{new Date().toLocaleString()}</span>
                  </div>
                  <div className="info-row">
                    <span>Expiration:</span>
                    <span>{new Date(Date.now() + 7 * 24 * 60 * 60 * 1000).toLocaleString()}</span>
                  </div>
                </div>
                <div className="channel-actions">
                  <button className="channel-action-btn">Top Up Balance</button>
                  <button className="channel-action-btn">Extend Duration</button>
                  <button className="channel-action-btn danger" onClick={closePaymentChannel}>Close Channel</button>
                </div>
              </div>
            )}
          </section>
        )}

        {activeTab === 'transactions' && (
          <section className="transactions-section">
            <h2>Transaction History</h2>
            <div className="filter-bar">
              <select>
                <option>All Transactions</option>
                <option>Successful</option>
                <option>Failed</option>
              </select>
              <button className="refresh-btn" onClick={() => setRequestCount(prev => prev + 1)}>
                Refresh
              </button>
            </div>

            {requestCount === 0 ? (
              <div className="empty-state">
                <p>No transactions yet. Start making RPC requests to see your transaction history.</p>
              </div>
            ) : (
              <div className="transactions-list">
                {Array.from({ length: Math.min(requestCount, 10) }).map((_, i) => (
                  <div key={i} className="transaction-item">
                    <div className="transaction-info">
                      <span className="transaction-method">eth_getBlockNumber</span>
                      <span className="transaction-time">{new Date(Date.now() - i * 60000).toLocaleString()}</span>
                    </div>
                    <div className="transaction-details">
                      <span className="transaction-fee">Fee: 0.001 ETH</span>
                      <span className={`transaction-status ${i % 9 === 0 ? 'failed' : 'success'}`}>
                        {i % 9 === 0 ? 'Failed' : 'Success'}
                      </span>
                    </div>
                  </div>
                ))}
              </div>
            )}
          </section>
        )}
      </main>

      <footer>
        <p>Decentralized RPC Network - A Protocol for Incentivized Blockchain Access</p>
        <div className="connection-status">
          <div className={`status-indicator ${selectedProvider ? 'connected' : 'disconnected'}`}></div>
          <span>{selectedProvider ? 'Connected' : 'Disconnected'}</span>
        </div>
      </footer>
    </div>
  )
}

export default App
