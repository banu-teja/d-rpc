import { useState, useEffect } from 'react'
import './App.css'
import { ethers } from 'ethers'

// API configuration
const API_ENDPOINT = import.meta.env.VITE_API_URL || 'http://localhost:8080';
const REGISTRY_CONTRACT_ADDRESS = import.meta.env.VITE_PROVIDER_REGISTRY_CONTRACT_ADDRESS || 'YOUR_PROVIDER_REGISTRY_CONTRACT_ADDRESS';
const STAKE_TOKEN_CONTRACT_ADDRESS = import.meta.env.VITE_STAKE_TOKEN_CONTRACT_ADDRESS || 'YOUR_STAKE_TOKEN_CONTRACT_ADDRESS';

// ABIs (Ideally, these would be imported from JSON files)
const providerRegistryABI = [
  // ... (Paste ProviderRegistry ABI here) ...
  {"type":"constructor","inputs":[{"name":"_stakeToken","type":"address","internalType":"contract IERC20"},{"name":"_minStake","type":"uint256","internalType":"uint256"}],"stateMutability":"nonpayable"},{"type":"function","name":"DEREGISTRATION_COOLDOWN","inputs":[],"outputs":[{"name":"","type":"uint256","internalType":"uint256"}],"stateMutability":"view"},{"type":"function","name":"calculateAndAllocateRewards","inputs":[{"name":"providersToReward","type":"address[]","internalType":"address[]"},{"name":"totalRewardAmount","type":"uint256","internalType":"uint256"}],"outputs":[],"stateMutability":"nonpayable"},{"type":"function","name":"claimRewards","inputs":[],"outputs":[],"stateMutability":"nonpayable"},{"type":"function","name":"depositStake","inputs":[{"name":"amount","type":"uint256","internalType":"uint256"}],"outputs":[],"stateMutability":"nonpayable"},{"type":"function","name":"deregister","inputs":[],"outputs":[],"stateMutability":"nonpayable"},{"type":"function","name":"fundRewardPool","inputs":[{"name":"amount","type":"uint256","internalType":"uint256"}],"outputs":[],"stateMutability":"nonpayable"},{"type":"function","name":"minStake","inputs":[],"outputs":[{"name":"","type":"uint256","internalType":"uint256"}],"stateMutability":"view"},{"type":"function","name":"owner","inputs":[],"outputs":[{"name":"","type":"address","internalType":"address"}],"stateMutability":"view"},{"type":"function","name":"providerRewards","inputs":[{"name":"","type":"address","internalType":"address"}],"outputs":[{"name":"","type":"uint256","internalType":"uint256"}],"stateMutability":"view"},{"type":"function","name":"providers","inputs":[{"name":"","type":"address","internalType":"address"}],"outputs":[{"name":"stake","type":"uint256","internalType":"uint256"},{"name":"qosScore","type":"uint256","internalType":"uint256"},{"name":"registered","type":"bool","internalType":"bool"},{"name":"registrationTime","type":"uint40","internalType":"uint40"},{"name":"deregistrationTime","type":"uint40","internalType":"uint40"},{"name":"endpointURL","type":"string","internalType":"string"}],"stateMutability":"view"},{"type":"function","name":"register","inputs":[],"outputs":[],"stateMutability":"nonpayable"},{"type":"function","name":"renounceOwnership","inputs":[],"outputs":[],"stateMutability":"nonpayable"},{"type":"function","name":"rewardPoolBalance","inputs":[],"outputs":[{"name":"","type":"uint256","internalType":"uint256"}],"stateMutability":"view"},{"type":"function","name":"setEndpointURL","inputs":[{"name":"_url","type":"string","internalType":"string"}],"outputs":[],"stateMutability":"nonpayable"},{"type":"function","name":"setMinStake","inputs":[{"name":"_minStake","type":"uint256","internalType":"uint256"}],"outputs":[],"stateMutability":"nonpayable"},{"type":"function","name":"slashProvider","inputs":[{"name":"provider_","type":"address","internalType":"address"},{"name":"amount","type":"uint256","internalType":"uint256"}],"outputs":[],"stateMutability":"nonpayable"},{"type":"function","name":"stakeToken","inputs":[],"outputs":[{"name":"","type":"address","internalType":"contract IERC20"}],"stateMutability":"view"},{"type":"function","name":"transferOwnership","inputs":[{"name":"newOwner","type":"address","internalType":"address"}],"outputs":[],"stateMutability":"nonpayable"},{"type":"function","name":"updateQoS","inputs":[{"name":"provider_","type":"address","internalType":"address"},{"name":"score","type":"uint256","internalType":"uint256"}],"outputs":[],"stateMutability":"nonpayable"},{"type":"function","name":"withdrawStake","inputs":[{"name":"amount","type":"uint256","internalType":"uint256"}],"outputs":[],"stateMutability":"nonpayable"},{"type":"event","name":"OwnershipTransferred","inputs":[{"name":"previousOwner","type":"address","indexed":true,"internalType":"address"},{"name":"newOwner","type":"address","indexed":true,"internalType":"address"}],"anonymous":false},{"type":"event","name":"ProviderDeregistered","inputs":[{"name":"provider","type":"address","indexed":true,"internalType":"address"}],"anonymous":false},{"type":"event","name":"ProviderRegistered","inputs":[{"name":"provider","type":"address","indexed":true,"internalType":"address"},{"name":"stake","type":"uint256","indexed":false,"internalType":"uint256"},{"name":"endpointURL","type":"string","indexed":false,"internalType":"string"}],"anonymous":false},{"type":"event","name":"ProviderSlashed","inputs":[{"name":"provider","type":"address","indexed":true,"internalType":"address"},{"name":"amount","type":"uint256","indexed":false,"internalType":"uint256"}],"anonymous":false},{"type":"event","name":"ProviderURLUpdated","inputs":[{"name":"provider","type":"address","indexed":true,"internalType":"address"},{"name":"newURL","type":"string","indexed":false,"internalType":"string"}],"anonymous":false},{\"type\":\"event\",\"name\":\"QoSUpdated\",\"inputs\":[{\"name\":\"provider\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"newScore\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RewardClaimed\",\"inputs\":[{\"name\":\"provider\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RewardPoolFunded\",\"inputs\":[{\"name\":\"funder\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RewardsAllocated\",\"inputs\":[{\"name\":\"allocator\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"totalAmount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"StakeDeposited\",\"inputs\":[{\"name\":\"provider\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"StakeWithdrawn\",\"inputs\":[{\"name\":\"provider\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"OwnableInvalidOwner\",\"inputs\":[{\"name\":\"owner\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"OwnableUnauthorizedAccount\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"ReentrancyGuardReentrantCall\",\"inputs\":[]}
];
const stakeTokenABI = [
  // ... (Paste StakeToken ABI here) ...
  {"type":"constructor","inputs":[],"stateMutability":"nonpayable"},{"type":"function","name":"allowance","inputs":[{"name":"owner","type":"address","internalType":"address"},{"name":"spender","type":"address","internalType":"address"}],"outputs":[{"name":"","type":"uint256","internalType":"uint256"}],"stateMutability":"view"},{"type":"function","name":"approve","inputs":[{"name":"spender","type":"address","internalType":"address"},{"name":"value","type":"uint256","internalType":"uint256"}],"outputs":[{"name":"","type":"bool","internalType":"bool"}],"stateMutability":"nonpayable"},{"type":"function","name":"balanceOf","inputs":[{"name":"account","type":"address","internalType":"address"}],"outputs":[{"name":"","type":"uint256","internalType":"uint256"}],"stateMutability":"view"},{"type":"function","name":"decimals","inputs":[],"outputs":[{"name":"","type":"uint8","internalType":"uint8"}],"stateMutability":"view"},{"type":"function","name":"name","inputs":[],"outputs":[{"name":"","type":"string","internalType":"string"}],"stateMutability":"view"},{"type":"function","name":"symbol","inputs":[],"outputs":[{"name":"","type":"string","internalType":"string"}],"stateMutability":"view"},{"type":"function","name":"totalSupply","inputs":[],"outputs":[{"name":"","type":"uint256","internalType":"uint256"}],"stateMutability":"view"},{"type":"function","name":"transfer","inputs":[{"name":"to","type":"address","internalType":"address"},{"name":"value","type":"uint256","internalType":"uint256"}],"outputs":[{"name":"","type":"bool","internalType":"bool"}],"stateMutability":"nonpayable"},{"type":"function","name":"transferFrom","inputs":[{"name":"from","type":"address","internalType":"address"},{"name":"to","type":"address","internalType":"address"},{"name":"value","type":"uint256","internalType":"uint256"}],"outputs":[{"name":"","type":"bool","internalType":"bool"}],"stateMutability":"nonpayable"},{"type":"event","name":"Approval","inputs":[{"name":"owner","type":"address","indexed":true,"internalType":"address"},{"name":"spender","type":"address","indexed":true,"internalType":"address"},{"name":"value","type":"uint256","indexed":false,"internalType":"uint256"}],"anonymous":false},{"type":"event","name":"Transfer","inputs":[{"name":"from","type":"address","indexed":true,"internalType":"address"},{"name":"to","type":"address","indexed":true,"internalType":"address"},{"name":"value","type":"uint256","indexed":false,"internalType":"uint256"}],"anonymous":false},{"type":"error","name":"ERC20InsufficientAllowance","inputs":[{"name":"spender","type":"address","internalType":"address"},{"name":"allowance","type":"uint256","internalType":"uint256"},{"name":"needed","type":"uint256","internalType":"uint256"}]},{\"type\":\"error\",\"name\":\"ERC20InsufficientBalance\",\"inputs\":[{\"name\":\"sender\",\"type\":\"address\"},{\"name\":\"balance\",\"type\":\"uint256\"},{\"name\":\"needed\",\"type\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"ERC20InvalidApprover\",\"inputs\":[{\"name\":\"approver\",\"type\":\"address\"}]},{\"type\":\"error\",\"name\":\"ERC20InvalidReceiver\",\"inputs\":[{\"name\":\"receiver\",\"type\":\"address\"}]},{\"type\":\"error\",\"name\":\"ERC20InvalidSender\",\"inputs\":[{\"name\":\"sender\",\"type\":\"address\"}]},{\"type\":\"error\",\"name\":\"ERC20InvalidSpender\",\"inputs\":[{\"name\":\"spender\",\"type\":\"address\"}]}
];

interface ProviderData {
  stake: string;
  qosScore: string;
  registered: boolean;
  registrationTime: number;
  deregistrationTime: number;
  endpointURL: string;
}

interface Provider {
=======
// API configuration
const API_ENDPOINT = import.meta.env.VITE_API_URL || 'http://localhost:8080';
const REGISTRY_CONTRACT_ADDRESS = import.meta.env.VITE_PROVIDER_REGISTRY_CONTRACT_ADDRESS || '0x5FbDB2315678afecb367f032d93F642f64180aa3'; // Default to Anvil Deployed Address
const STAKE_TOKEN_CONTRACT_ADDRESS = import.meta.env.VITE_STAKE_TOKEN_CONTRACT_ADDRESS || '0xe7f1725E7734CE288F8367e1Bb143E90bb3F0512'; // Default to Anvil Deployed Address

// ABIs
const providerRegistryABI = JSON.parse('[{"type":"constructor","inputs":[{"name":"_stakeToken","type":"address","internalType":"contract IERC20"},{"name":"_minStake","type":"uint256","internalType":"uint256"}],"stateMutability":"nonpayable"},{"type":"function","name":"DEREGISTRATION_COOLDOWN","inputs":[],"outputs":[{"name":"","type":"uint256","internalType":"uint256"}],"stateMutability":"view"},{"type":"function","name":"calculateAndAllocateRewards","inputs":[{"name":"providersToReward","type":"address[]","internalType":"address[]"},{"name":"totalRewardAmount","type":"uint256","internalType":"uint256"}],"outputs":[],"stateMutability":"nonpayable"},{"type":"function","name":"claimRewards","inputs":[],"outputs":[],"stateMutability":"nonpayable"},{"type":"function","name":"depositStake","inputs":[{"name":"amount","type":"uint256","internalType":"uint256"}],"outputs":[],"stateMutability":"nonpayable"},{"type":"function","name":"deregister","inputs":[],"outputs":[],"stateMutability":"nonpayable"},{"type":"function","name":"fundRewardPool","inputs":[{"name":"amount","type":"uint256","internalType":"uint256"}],"outputs":[],"stateMutability":"nonpayable"},{"type":"function","name":"minStake","inputs":[],"outputs":[{"name":"","type":"uint256","internalType":"uint256"}],"stateMutability":"view"},{"type":"function","name":"owner","inputs":[],"outputs":[{"name":"","type":"address","internalType":"address"}],"stateMutability":"view"},{"type":"function","name":"providerRewards","inputs":[{"name":"","type":"address","internalType":"address"}],"outputs":[{"name":"","type":"uint256","internalType":"uint256"}],"stateMutability":"view"},{"type":"function","name":"providers","inputs":[{"name":"","type":"address","internalType":"address"}],"outputs":[{"name":"stake","type":"uint256","internalType":"uint256"},{"name":"qosScore","type":"uint256","internalType":"uint256"},{"name":"registered","type":"bool","internalType":"bool"},{"name":"registrationTime","type":"uint40","internalType":"uint40"},{"name":"deregistrationTime","type":"uint40","internalType":"uint40"},{"name":"endpointURL","type":"string","internalType":"string"}],"stateMutability":"view"},{"type":"function","name":"register","inputs":[],"outputs":[],"stateMutability":"nonpayable"},{"type":"function","name":"renounceOwnership","inputs":[],"outputs":[],"stateMutability":"nonpayable"},{"type":"function","name":"rewardPoolBalance","inputs":[],"outputs":[{"name":"","type":"uint256","internalType":"uint256"}],"stateMutability":"view"},{"type":"function","name":"setEndpointURL","inputs":[{"name":"_url","type":"string","internalType":"string"}],"outputs":[],"stateMutability":"nonpayable"},{"type":"function","name":"setMinStake","inputs":[{"name":"_minStake","type":"uint256","internalType":"uint256"}],"outputs":[],"stateMutability":"nonpayable"},{"type":"function","name":"slashProvider","inputs":[{"name":"provider_","type":"address","internalType":"address"},{"name":"amount","type":"uint256","internalType":"uint256"}],"outputs":[],"stateMutability":"nonpayable"},{"type":"function","name":"stakeToken","inputs":[],"outputs":[{"name":"","type":"address","internalType":"contract IERC20"}],"stateMutability":"view"},{"type":"function","name":"transferOwnership","inputs":[{"name":"newOwner","type":"address","internalType":"address"}],"outputs":[],"stateMutability":"nonpayable"},{"type":"function","name":"updateQoS","inputs":[{"name":"provider_","type":"address","internalType":"address"},{"name":"score","type":"uint256","internalType":"uint256"}],"outputs":[],"stateMutability":"nonpayable"},{"type":"function","name":"withdrawStake","inputs":[{"name":"amount","type":"uint256","internalType":"uint256"}],"outputs":[],"stateMutability":"nonpayable"},{"type":"event","name":"OwnershipTransferred","inputs":[{"name":"previousOwner","type":"address","indexed":true,"internalType":"address"},{"name":"newOwner","type":"address","indexed":true,"internalType":"address"}],"anonymous":false},{"type":"event","name":"ProviderDeregistered","inputs":[{"name":"provider","type":"address","indexed":true,"internalType":"address"}],"anonymous":false},{"type":"event","name":"ProviderRegistered","inputs":[{"name":"provider","type":"address","indexed":true,"internalType":"address"},{"name":"stake","type":"uint256","indexed":false,"internalType":"uint256"},{"name":"endpointURL","type":"string","indexed":false,"internalType":"string"}],"anonymous":false},{"type":"event","name":"ProviderSlashed","inputs":[{"name":"provider","type":"address","indexed":true,"internalType":"address"},{"name":"amount","type":"uint256","indexed":false,"internalType":"uint256"}],"anonymous":false},{"type":"event","name":"ProviderURLUpdated","inputs":[{"name":"provider","type":"address","indexed":true,"internalType":"address"},{"name":"newURL","type":"string","indexed":false,"internalType":"string"}],"anonymous":false},{\"type\":\"event\",\"name\":\"QoSUpdated\",\"inputs\":[{\"name\":\"provider\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"newScore\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RewardClaimed\",\"inputs\":[{\"name\":\"provider\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RewardPoolFunded\",\"inputs\":[{\"name\":\"funder\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RewardsAllocated\",\"inputs\":[{\"name\":\"allocator\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"totalAmount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"StakeDeposited\",\"inputs\":[{\"name\":\"provider\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"StakeWithdrawn\",\"inputs\":[{\"name\":\"provider\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"OwnableInvalidOwner\",\"inputs\":[{\"name\":\"owner\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"OwnableUnauthorizedAccount\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"ReentrancyGuardReentrantCall\",\"inputs\":[]}]');
const stakeTokenABI = JSON.parse('[{"type":"constructor","inputs":[],"stateMutability":"nonpayable"},{"type":"function","name":"allowance","inputs":[{"name":"owner","type":"address","internalType":"address"},{"name":"spender","type":"address","internalType":"address"}],"outputs":[{"name":"","type":"uint256","internalType":"uint256"}],"stateMutability":"view"},{"type":"function","name":"approve","inputs":[{"name":"spender","type":"address","internalType":"address"},{"name":"value","type":"uint256","internalType":"uint256"}],"outputs":[{"name":"","type":"bool","internalType":"bool"}],"stateMutability":"nonpayable"},{"type":"function","name":"balanceOf","inputs":[{"name":"account","type":"address","internalType":"address"}],"outputs":[{"name":"","type":"uint256","internalType":"uint256"}],"stateMutability":"view"},{"type":"function","name":"decimals","inputs":[],"outputs":[{"name":"","type":"uint8","internalType":"uint8"}],"stateMutability":"view"},{"type":"function","name":"name","inputs":[],"outputs":[{"name":"","type":"string","internalType":"string"}],"stateMutability":"view"},{"type":"function","name":"symbol","inputs":[],"outputs":[{"name":"","type":"string","internalType":"string"}],"stateMutability":"view"},{"type":"function","name":"totalSupply","inputs":[],"outputs":[{"name":"","type":"uint256","internalType":"uint256"}],"stateMutability":"view"},{"type":"function","name":"transfer","inputs":[{"name":"to","type":"address","internalType":"address"},{"name":"value","type":"uint256","internalType":"uint256"}],"outputs":[{"name":"","type":"bool","internalType":"bool"}],"stateMutability":"nonpayable"},{"type":"function","name":"transferFrom","inputs":[{"name":"from","type":"address","internalType":"address"},{"name":"to","type":"address","internalType":"address"},{"name":"value","type":"uint256","internalType":"uint256"}],"outputs":[{"name":"","type":"bool","internalType":"bool"}],"stateMutability":"nonpayable"},{"type":"event","name":"Approval","inputs":[{"name":"owner","type":"address","indexed":true,"internalType":"address"},{"name":"spender","type":"address","indexed":true,"internalType":"address"},{"name":"value","type":"uint256","indexed":false,"internalType":"uint256"}],"anonymous":false},{"type":"event","name":"Transfer","inputs":[{"name":"from","type":"address","indexed":true,"internalType":"address"},{"name":"to","type":"address","indexed":true,"internalType":"address"},{"name":"value","type":"uint256","indexed":false,"internalType":"uint256"}],"anonymous":false},{"type":"error","name":"ERC20InsufficientAllowance","inputs":[{"name":"spender","type":"address","internalType":"address"},{"name":"allowance","type":"uint256","internalType":"uint256"},{"name":"needed","type":"uint256","internalType":"uint256"}]},{\"type\":\"error\",\"name\":\"ERC20InsufficientBalance\",\"inputs\":[{\"name\":\"sender\",\"type\":\"address\"},{\"name\":\"balance\",\"type\":\"uint256\"},{\"name\":\"needed\",\"type\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"ERC20InvalidApprover\",\"inputs\":[{\"name\":\"approver\",\"type\":\"address\"}]},{\"type\":\"error\",\"name\":\"ERC20InvalidReceiver\",\"inputs\":[{\"name\":\"receiver\",\"type\":\"address\"}]},{\"type\":\"error\",\"name\":\"ERC20InvalidSender\",\"inputs\":[{\"name\":\"sender\",\"type\":\"address\"}]},{\"type\":\"error\",\"name\":\"ERC20InvalidSpender\",\"inputs\":[{\"name\":\"spender\",\"type\":\"address\"}]}]');

interface ProviderContractData {
  stake: bigint;
  qosScore: bigint;
  registered: boolean;
  registrationTime: bigint; // Solidity uint40 is best represented as bigint in JS
  deregistrationTime: bigint;
  endpointURL: string;
}

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
  const [signer, setSigner] = useState<ethers.Signer | null>(null)
  const [account, setAccount] = useState<string | null>(null)
  const [ethersProvider, setEthersProvider] = useState<ethers.BrowserProvider | null>(null)
  const [providerRegistryContract, setProviderRegistryContract] = useState<ethers.Contract | null>(null)
  const [stakeTokenContract, setStakeTokenContract] = useState<ethers.Contract | null>(null)
  const [operatorData, setOperatorData] = useState<ProviderContractData | null>(null)
  const [operatorRewards, setOperatorRewards] = useState<bigint>(BigInt(0))
  const [stakeAmount, setStakeAmount] = useState('100') // Default STK, assuming 18 decimals
  const [endpointUrlInput, setEndpointUrlInput] = useState('')
  const [minStakeRequired, setMinStakeRequired] = useState<bigint>(BigInt(0))
  const [txMessage, setTxMessage] = useState<string | null>(null)
  const [txError, setTxError] = useState<string | null>(null)

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

  // Modified to accept an optional provider (used after wallet connection)
  const fetchNetworkInfo = async (providerInstance?: ethers.Provider) => {
    try {
      // Use the connected provider if available, otherwise fall back to gateway RPC
      const provider = providerInstance || (ethersProvider ?? new ethers.JsonRpcProvider(`${rpcEndpoint}`))
      const network = await provider.getNetwork()
      setNetworkInfo({
        // Use short name for anvil for brevity
        name: network.name === 'unknown' ? 'Anvil' : network.name,
        chainId: Number(network.chainId)
      })
    } catch (err) {
      console.error("Error fetching network info:", err)
    }
  }

  const connectWallet = async () => {
    if (typeof window.ethereum !== 'undefined') {
      try {
        const provider = new ethers.BrowserProvider(window.ethereum)
        await provider.send("eth_requestAccounts", []);
        const currentSigner = await provider.getSigner();
        const currentAccount = await currentSigner.getAddress();

        setEthersProvider(provider);
        setSigner(currentSigner);
        setAccount(currentAccount);

        const registry = new ethers.Contract(REGISTRY_CONTRACT_ADDRESS, providerRegistryABI, currentSigner);
        setProviderRegistryContract(registry);

        const tokenAddr = await registry.stakeToken();
        const token = new ethers.Contract(tokenAddr, stakeTokenABI, currentSigner);
        setStakeTokenContract(token);

        fetchMinStake(registry);
        fetchOperatorData(registry, currentAccount);
        fetchOperatorRewards(registry, currentAccount);

        // Fetch initial network info after connecting
        fetchNetworkInfo(provider); // Pass the connected provider

        // Listen for account changes
        window.ethereum.on('accountsChanged', (accounts: string[]) => {
          if (accounts.length > 0) {
            setAccount(accounts[0]);
            // Re-fetch signer if needed, though BrowserProvider usually handles this
            provider.getSigner().then(newSigner => {
              setSigner(newSigner);
              const newRegistry = new ethers.Contract(REGISTRY_CONTRACT_ADDRESS, providerRegistryABI, newSigner);
              setProviderRegistryContract(newRegistry);
              newRegistry.stakeToken().then((tokenAddr: string) => {
                const newToken = new ethers.Contract(tokenAddr, stakeTokenABI, newSigner);
                setStakeTokenContract(newToken);
              });
              fetchOperatorData(newRegistry, accounts[0]);
              fetchOperatorRewards(newRegistry, accounts[0]);
            });
          } else {
            setAccount(null);
            setSigner(null);
            setEthersProvider(null);
            setProviderRegistryContract(null);
            setStakeTokenContract(null);
            setOperatorData(null);
            setOperatorRewards(BigInt(0));
          }
        });

        // Listen for chain changes
        window.ethereum.on('chainChanged', (_chainId: string) => {
          // Reload the page to ensure everything is re-initialized correctly
          window.location.reload();
        });

      } catch (err) {
        console.error("Wallet connection error:", err);
        setError(`Failed to connect wallet: ${err instanceof Error ? err.message : String(err)}`);
        setAccount(null);
        setSigner(null);
        setEthersProvider(null);
      }
    } else {
      setError('MetaMask (or another Ethereum wallet) is not installed.');
    }
  };

  const fetchMinStake = async (registry: ethers.Contract) => {
    try {
      const min = await registry.minStake();
      setMinStakeRequired(min);
    } catch (err) {
      console.error("Error fetching min stake:", err);
      setTxError("Failed to fetch min stake.")
    }
  }

  const fetchOperatorData = async (registry: ethers.Contract, operatorAddress: string) => {
    if (!registry || !operatorAddress) return;
    try {
      const data = await registry.providers(operatorAddress);
      setOperatorData({
        stake: data.stake,
        qosScore: data.qosScore,
        registered: data.registered,
        registrationTime: data.registrationTime,
        deregistrationTime: data.deregistrationTime,
        endpointURL: data.endpointURL,
      });
      setEndpointUrlInput(data.endpointURL || '');
    } catch (err) {
      console.error("Error fetching operator data:", err);
      setOperatorData(null);
      setTxError("Failed to fetch operator data.")
    }
  };

  const fetchOperatorRewards = async (registry: ethers.Contract, operatorAddress: string) => {
    if (!registry || !operatorAddress) return;
    try {
      const rewards = await registry.providerRewards(operatorAddress);
      setOperatorRewards(rewards);
    } catch (err) {
      console.error("Error fetching operator rewards:", err);
      setOperatorRewards(BigInt(0));
      setTxError("Failed to fetch operator rewards.")
    }
  };

  const handleTransaction = async (txPromise: Promise<any>, successMessage: string) => {
    setTxMessage('Processing transaction...');
    setTxError(null);
    try {
      const tx = await txPromise;
      setTxMessage('Waiting for confirmation...');
      await tx.wait();
      setTxMessage(successMessage);
      // Refresh relevant data
      if (providerRegistryContract && account) {
        fetchOperatorData(providerRegistryContract, account);
        fetchOperatorRewards(providerRegistryContract, account);
        fetchMinStake(providerRegistryContract);
      }
    } catch (err: any) {
      console.error("Transaction error:", err);
      setTxError(err.reason || err.message || 'Transaction failed.');
      setTxMessage(null);
    }
  };

  const handleDepositStake = async () => {
    if (!stakeTokenContract || !providerRegistryContract || !stakeAmount) return;
    const amount = ethers.parseUnits(stakeAmount, 18); // Assuming 18 decimals for STK
    await handleTransaction(
      stakeTokenContract.approve(REGISTRY_CONTRACT_ADDRESS, amount).then(async (tx: any) => {
        await tx.wait();
        setTxMessage('Approval successful, now depositing stake...');
        return providerRegistryContract.depositStake(amount);
      }),
      'Stake deposited successfully!'
    );
  };

  const handleSetEndpoint = async () => {
    if (!providerRegistryContract || !endpointUrlInput) return;
    await handleTransaction(
      providerRegistryContract.setEndpointURL(endpointUrlInput),
      'Endpoint URL updated successfully!'
    );
  };

 const handleRegister = async () => {
    if (!providerRegistryContract) return;
    await handleTransaction(
      providerRegistryContract.register(),
      'Provider registered successfully!'
    );
  };

  const handleDeregister = async () => {
    if (!providerRegistryContract) return;
    await handleTransaction(
      providerRegistryContract.deregister(),
      'Provider deregistered successfully! Cooldown period applies for withdrawal.'
    );
  };

  const handleWithdrawStake = async () => {
    if (!providerRegistryContract || !stakeAmount) return;
    const amount = ethers.parseUnits(stakeAmount, 18);
    await handleTransaction(
      providerRegistryContract.withdrawStake(amount),
      'Stake withdrawn successfully!'
    );
  };

  const handleClaimRewards = async () => {
    if (!providerRegistryContract) return;
    await handleTransaction(
      providerRegistryContract.claimRewards(),
      'Rewards claimed successfully!'
    );
  };

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
        <div className="header-content">
            <h1>Decentralized RPC Network</h1>
            <p>A network of incentivized RPC providers</p>
        </div>
        <div className="wallet-connect">
          {account ? (
            <div className="account-info">
              <span>Connected: {`${account.substring(0, 6)}...${account.substring(account.length - 4)}`}</span>
              {networkInfo && <span>Network: {networkInfo.name} ({networkInfo.chainId})</span>}
            </div>
          ) : (
            <button onClick={connectWallet} className="connect-wallet-btn">Connect Wallet</button>
          )}
        </div>

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
            <li className={activeTab === 'operator' ? 'active' : ''}>
              <button 
                onClick={() => {
                  setActiveTab('operator');
                  if(providerRegistryContract && account) {
                    fetchOperatorData(providerRegistryContract, account);
                    fetchOperatorRewards(providerRegistryContract, account);
                    fetchMinStake(providerRegistryContract);
                  }
                }}
                disabled={!account}
              >
                Operator Zone
              </button>
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

        {activeTab === 'operator' && account && signer && (
          <section className="operator-section">
            <h2>Node Operator Zone</h2>
            <p>Manage your node registration, stake, and rewards.</p>

            {/* Placeholder for Operator UI elements */}
            <div className="operator-actions">
              {txMessage && <div className="tx-message success-message">{txMessage}</div>}
              {txError && <div className="tx-message error-message">{txError}</div>}

              <div className="operator-info card">
                <h3>Your Node Status</h3>
                {operatorData ? (
                  <>
                    <p><strong>Registered:</strong> {operatorData.registered ? 'Yes' : 'No'}</p>
                    <p><strong>Current Stake:</strong> {ethers.formatUnits(operatorData.stake, 18)} STK</p>
                    <p><strong>Min Stake Required:</strong> {ethers.formatUnits(minStakeRequired, 18)} STK</p>
                    <p><strong>QoS Score:</strong> {operatorData.qosScore.toString()}</p>
                    <p><strong>Endpoint URL:</strong> {operatorData.endpointURL || 'Not Set'}</p>
                    <p><strong>Registration Time:</strong> {operatorData.registrationTime > 0 ? new Date(Number(operatorData.registrationTime) * 1000).toLocaleString() : 'N/A'}</p>
                    <p><strong>Deregistration Time:</strong> {operatorData.deregistrationTime > 0 ? new Date(Number(operatorData.deregistrationTime) * 1000).toLocaleString() : 'N/A'}</p>
                  </>
                ) : (
                  <p>Loading operator data...</p>
                )}
              </div>

              <div className="operator-rewards card">
                <h3>Your Rewards</h3>
                <p><strong>Claimable Rewards:</strong> {ethers.formatUnits(operatorRewards, 18)} STK</p>
                <button onClick={handleClaimRewards} disabled={!providerRegistryContract || operatorRewards === BigInt(0)}>Claim Rewards</button>
              </div>

              <div className="operator-stake card">
                <h3>Manage Stake</h3>
                <div className="form-group">
                  <label>Amount (STK):</label>
                  <input type="number" value={stakeAmount} onChange={(e) => setStakeAmount(e.target.value)} />
                </div>
                <button onClick={handleDepositStake} disabled={!stakeTokenContract || !providerRegistryContract}>Deposit Stake</button>
                <p className="info-text">To withdraw, deregister first and wait for cooldown.</p>
                <button onClick={handleWithdrawStake} disabled={!providerRegistryContract || !(operatorData && !operatorData.registered && operatorData.stake > BigInt(0)) } >Withdraw Stake</button>
              </div>

              <div className="operator-registration card">
                <h3>Manage Registration</h3>
                <div className="form-group">
                  <label>Endpoint URL (e.g., http://your-node-ip:8545):</label>
                  <input type="text" value={endpointUrlInput} onChange={(e) => setEndpointUrlInput(e.target.value)} placeholder="http://your-node-ip:8545" />
                </div>
                <button onClick={handleSetEndpoint} disabled={!providerRegistryContract}>Set/Update Endpoint URL</button>
                
                {operatorData && !operatorData.registered ? (
                  <button onClick={handleRegister} disabled={!providerRegistryContract || !operatorData?.endpointURL || operatorData?.stake < minStakeRequired}>
                    Register Node
                  </button>
                ) : (
                  <button onClick={handleDeregister} disabled={!providerRegistryContract || !(operatorData && operatorData.registered)}>
                    Deregister Node
                  </button>
                )}
                {!(operatorData && operatorData.registered) && operatorData?.stake < minStakeRequired && <p className="error-text">Insufficient stake to register.</p>}
                 {!(operatorData && operatorData.registered) && !operatorData?.endpointURL && <p className="error-text">Endpoint URL must be set to register.</p>}
              </div>

            </div>
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
