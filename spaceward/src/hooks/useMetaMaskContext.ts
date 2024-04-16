import { useContext } from "react";
import { MetaMaskContext } from "../context/MetaMaskContext";

/**
 * Utility hook to consume the MetaMask context.
 *
 * @returns The MetaMask context.
 */
export default function useMetaMaskContext() {
    return useContext(MetaMaskContext);
}
