class Networks {
    // EVM
    static ETH_MAINNET = 1
    static BSC_MAINNET = 56  
    static POLYGON_MAINNET = 137
    static AVALANCHE_MAINNET = 43114
    static ARBITRUM_MAINNET = 42161
    static OPTIMISM_MAINNET = 10
    static FANTOM_MAINNET = 250
    static CRONOS_MAINNET = 25
    static GNOSIS_MAINNET = 100
    static KLAYTN_MAINNET = 8217
    static AURORA_MAINNET = 1313161554
    static CELO_MAINNET = 42220
    static MOONBEAM_MAINNET = 1284
    static MOONRIVER_MAINNET = 1285
    static METIS_MAINNET = 1088
    static BASE_MAINNET = 8453
    static LINEA_MAINNET = 59144
    static ZKSYNC_MAINNET = 324
    static ZKEVM_MAINNET = 1101
    static MANTLE_MAINNET = 5000
    static SCROLL_MAINNET = 534352
    static KAVA_MAINNET = 2222
    static HARMONY_MAINNET = 1666600000
    static FUSE_MAINNET = 122
    static OKX_MAINNET = 66
    static CORE_MAINNET = 1116
    static PULSE_MAINNET = 369
    static EOS_EVM_MAINNET = 17777
    static TELOS_MAINNET = 40
    static BRISE_MAINNET = 32520
    static CONFLUX_MAINNET = 1030
    static OASIS_MAINNET = 42262
    static ELASTOS_MAINNET = 20
    static IOTEX_MAINNET = 4689
    static SYSCOIN_MAINNET = 57
    static CANTO_MAINNET = 7700
    static BOBA_MAINNET = 288
    static VELAS_MAINNET = 106
    static ASTAR_MAINNET = 592
    static THETA_MAINNET = 361
    static RONIN_MAINNET = 2020
    static CHILIZ_MAINNET = 88
    static LUKSO_MAINNET = 42
    static COINEX_MAINNET = 52
    static WEMIX_MAINNET = 1111
    static THUNDERCORE_MAINNET = 108
    static ULTRON_MAINNET = 1231
    static OASYS_MAINNET = 248
    static SHIDEN_MAINNET = 336
    static DOGECHAIN_MAINNET = 2000
    static TENET_MAINNET = 1559
    static NOVA_MAINNET = 87
    static EVMOS_MAINNET = 9001
    static PALM_MAINNET = 11297108109
 
    // NOT-EVM
    static SOLANA_MAINNET = "mainnet-beta"
    static NEAR_MAINNET = "mainnet"
    static APTOS_MAINNET = 1
    static SUI_MAINNET = "mainnet"
    static CARDANO_MAINNET = "mainnet"
    static POLKADOT_MAINNET = "polkadot"
    static COSMOS_MAINNET = "cosmoshub-4"
    static TON_MAINNET = "-239"
    static TRON_MAINNET = "0x2b6653dc"
    static BEAM_MAINNET = "mainnet"
    static STARKNET_MAINNET = "SN_MAIN"
    static HEDERA_MAINNET = "mainnet"
    static KADENA_MAINNET = "mainnet"
    static MINA_MAINNET = "mainnet"
    static ALGORAND_MAINNET = "mainnet"
    static FLOW_MAINNET = "mainnet"
    static WAVES_MAINNET = "mainnet"
    static RIPPLE_MAINNET = "mainnet"
    static STELLAR_MAINNET = "mainnet"
    static TEZOS_MAINNET = "mainnet"
    static IOTA_MAINNET = "mainnet"
    static ZILLIQA_MAINNET = "mainnet"
    static NEO_MAINNET = "mainnet"
    static ERGO_MAINNET = "mainnet"
    static INTERNET_COMPUTER = "mainnet"
    static CASPER_MAINNET = "mainnet"
    static ARK_MAINNET = "mainnet"
    static TERRA_MAINNET = "phoenix-1"
    static THORCHAIN_MAINNET = "mainnet"
    static INJECTIVE_MAINNET = "mainnet"
    static SEI_MAINNET = "pacific-1"
    static CELESTIA_MAINNET = "mainnet"
    static SXN_MAINNET = "mainnet"
 
    static getNetworkName(chainId) {
        const networkNames = {
            // EVM 
            [Networks.ETH_MAINNET]: "Ethereum Network",
            [Networks.BSC_MAINNET]: "BNB Smart Chain",
            [Networks.POLYGON_MAINNET]: "Polygon Network",
            [Networks.AVALANCHE_MAINNET]: "Avalanche Network",
            [Networks.ARBITRUM_MAINNET]: "Arbitrum Network",
            [Networks.OPTIMISM_MAINNET]: "Optimism Network",
            [Networks.FANTOM_MAINNET]: "Fantom Network",
            [Networks.CRONOS_MAINNET]: "Cronos Network",
            [Networks.GNOSIS_MAINNET]: "Gnosis Network",
            [Networks.KLAYTN_MAINNET]: "Klaytn Network",
            [Networks.AURORA_MAINNET]: "Aurora Network",
            [Networks.CELO_MAINNET]: "Celo Network",
            [Networks.MOONBEAM_MAINNET]: "Moonbeam Network",
            [Networks.MOONRIVER_MAINNET]: "Moonriver Network",
            [Networks.METIS_MAINNET]: "Metis Network",
            [Networks.BASE_MAINNET]: "Base Network",
            [Networks.LINEA_MAINNET]: "Linea Network",
            [Networks.ZKSYNC_MAINNET]: "zkSync Network",
            [Networks.ZKEVM_MAINNET]: "zkEVM Network",
            [Networks.MANTLE_MAINNET]: "Mantle Network",
            [Networks.SCROLL_MAINNET]: "Scroll Network",
            [Networks.KAVA_MAINNET]: "Kava Network",
            [Networks.HARMONY_MAINNET]: "Harmony Network",
            [Networks.FUSE_MAINNET]: "Fuse Network",
            [Networks.OKX_MAINNET]: "OKX Network",
            [Networks.CORE_MAINNET]: "Core Network",
            [Networks.PULSE_MAINNET]: "PulseChain Network",
            [Networks.EOS_EVM_MAINNET]: "EOS EVM Network",
            [Networks.TELOS_MAINNET]: "Telos Network",
            [Networks.BRISE_MAINNET]: "Bitgert Network",
            [Networks.CONFLUX_MAINNET]: "Conflux Network",
            [Networks.OASIS_MAINNET]: "Oasis Network",
            [Networks.ELASTOS_MAINNET]: "Elastos Network",
            [Networks.IOTEX_MAINNET]: "IoTeX Network",
            [Networks.SYSCOIN_MAINNET]: "Syscoin Network",
            [Networks.CANTO_MAINNET]: "Canto Network",
            [Networks.BOBA_MAINNET]: "Boba Network",
            [Networks.VELAS_MAINNET]: "Velas Network",
            [Networks.ASTAR_MAINNET]: "Astar Network",
            [Networks.THETA_MAINNET]: "Theta Network",
            [Networks.RONIN_MAINNET]: "Ronin Network",
            [Networks.CHILIZ_MAINNET]: "Chiliz Network",
            [Networks.LUKSO_MAINNET]: "Lukso Network",
            [Networks.COINEX_MAINNET]: "CoinEx Network",
            [Networks.WEMIX_MAINNET]: "Wemix Network",
            [Networks.THUNDERCORE_MAINNET]: "ThunderCore Network",
            [Networks.ULTRON_MAINNET]: "Ultron Network",
            [Networks.OASYS_MAINNET]: "Oasys Network",
            [Networks.SHIDEN_MAINNET]: "Shiden Network",
            [Networks.DOGECHAIN_MAINNET]: "Dogechain Network",
            [Networks.TENET_MAINNET]: "Tenet Network",
            [Networks.NOVA_MAINNET]: "Nova Network",
            [Networks.EVMOS_MAINNET]: "Evmos Network",
            [Networks.PALM_MAINNET]: "Palm Network",
 
            // not-EVM 
            [Networks.SOLANA_MAINNET]: "Solana Network",
            [Networks.NEAR_MAINNET]: "Near Network",
            [Networks.APTOS_MAINNET]: "Aptos Network",
            [Networks.SUI_MAINNET]: "Sui Network",
            [Networks.CARDANO_MAINNET]: "Cardano Network",
            [Networks.POLKADOT_MAINNET]: "Polkadot Network",
            [Networks.COSMOS_MAINNET]: "Cosmos Network",
            [Networks.TON_MAINNET]: "TON Network",
            [Networks.TRON_MAINNET]: "TRON Network",
            [Networks.BEAM_MAINNET]: "Beam Network",
            [Networks.STARKNET_MAINNET]: "Starknet Network",
            [Networks.HEDERA_MAINNET]: "Hedera Network",
            [Networks.KADENA_MAINNET]: "Kadena Network",
            [Networks.MINA_MAINNET]: "Mina Network",
            [Networks.ALGORAND_MAINNET]: "Algorand Network",
            [Networks.FLOW_MAINNET]: "Flow Network",
            [Networks.WAVES_MAINNET]: "Waves Network",
            [Networks.RIPPLE_MAINNET]: "Ripple Network",
            [Networks.STELLAR_MAINNET]: "Stellar Network",
            [Networks.TEZOS_MAINNET]: "Tezos Network",
            [Networks.IOTA_MAINNET]: "IOTA Network",
            [Networks.ZILLIQA_MAINNET]: "Zilliqa Network",
            [Networks.NEO_MAINNET]: "NEO Network",
            [Networks.ERGO_MAINNET]: "Ergo Network",
            [Networks.INTERNET_COMPUTER]: "Internet Computer Network",
            [Networks.CASPER_MAINNET]: "Casper Network",
            [Networks.ARK_MAINNET]: "Ark Network",
            [Networks.TERRA_MAINNET]: "Terra Network",
            [Networks.THORCHAIN_MAINNET]: "THORChain Network",
            [Networks.INJECTIVE_MAINNET]: "Injective Network",
            [Networks.SEI_MAINNET]: "Sei Network",
            [Networks.CELESTIA_MAINNET]: "Celestia Network",
            [Networks.SXN_MAINNET]: "SXNetwork"
        }
        
        return networkNames[chainId] || undefined
    }
 }