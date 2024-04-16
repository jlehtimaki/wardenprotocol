import Long from "long";
import { Skeleton } from "@/components/ui/skeleton";
import AddressAvatar from "@/components/AddressAvatar";
import { ethers } from "ethers";
import { useQuery } from "@tanstack/react-query";
import { Button } from "@/components/ui/button";
import { MoveUpRight, KeyIcon } from "lucide-react";
import { Link } from "react-router-dom";
import { useSpaceId } from "@/hooks/useSpaceId";
import { NewKeyButton } from "@/features/keys";
import { ReceiveAssetButton } from "@/features/assets";
import { Copy } from "@/components/ui/copy";
import { useQueryHooks } from "@/hooks/useClient";
import { AddressType } from "@wardenprotocol/wardjs/dist/codegen/warden/warden/v1beta2/key";
import { PageRequest } from "@wardenprotocol/wardjs/dist/codegen/cosmos/base/query/v1beta1/pagination";

const url = "https://rpc2.sepolia.org";
const provider = new ethers.JsonRpcProvider(url);

async function getEthBalance(address: string) {
	const balance = await provider.getBalance(address);
	return balance;
}

export function HomeAssets() {
	const { spaceId } = useSpaceId();
	const { useKeysBySpaceId } = useQueryHooks();
	const query = useKeysBySpaceId({
		request: {
			spaceId: Long.fromString(spaceId || ""),
			deriveAddresses: [
				AddressType.ADDRESS_TYPE_ETHEREUM,
				AddressType.ADDRESS_TYPE_OSMOSIS,
			],
			pagination: PageRequest.fromPartial({
				limit: Long.fromInt(10),
			}),
		},
		options: {
			enabled: !!spaceId,
		},
	});

	if (query.status === "loading") {
		return (
			<div className="flex h-60 flex-col space-y-1 items-center place-content-center">
				<span className="text-muted-foreground">Loading...</span>
			</div>
		);
	}

	if (query.data?.keys?.length === 0) {
		return (
			<div className="flex h-60 flex-col space-y-1 items-center place-content-center">
				<KeyIcon className="h-10 w-10" />
				<span className="pt-4">No keys found in this space</span>
				<span className="text-muted-foreground text-sm pb-4">
					Add a key to start receiving assets
				</span>
				<NewKeyButton />
			</div>
		);
	}

	return (
		<>
			{query.data?.keys?.map((key) => (
				<div className="flex flex-col border m-4 rounded-lg">
					<div>
						<div className="space-y-3">
							{key.addresses?.map((addr) => {
								return (
									<Address
										address={addr.address || ""}
										type={addr.type}
										keyId={key.key.id}
									/>
								);
							})}
						</div>
					</div>
				</div>
			))}
		</>
	);
}

function Address({
	address,
	type,
	keyId,
}: {
	address: string;
	type: AddressType;
	keyId: Long;
}) {
	if (type != AddressType.ADDRESS_TYPE_ETHEREUM) {
		return null;
	}

	return (
		<div className="">
			<div className="py-1 hover:no-underline border-0 font-normal font-sans">
				<div className="flex flex-col md:flex-row justify-between w-full mr-4 gap-4 p-4">
					<div className="flex flex-row items-center gap-4">
						<AddressAvatar seed={address} />
						<div className="font-sans flex flex-col text-left">
							<span className="text-muted-foreground text-xs">
								Wallet Address
							</span>
							<div className="flex flex-row gap-2 items-center">
								<Copy value={address} split={true} />
							</div>
						</div>
					</div>
					<div className="flex flex-row items-center">
						<ReceiveAssetButton address={address} />
					</div>
				</div>
			</div>
			<div>
				<div className="space-y-4">
					<Sepolia address={address} keyId={keyId} />
				</div>
			</div>
		</div>
	);
}

function Sepolia({ address, keyId }: { address: string; keyId: Long }) {
	const query = useQuery({
		queryKey: ["eth-balance", address],
		queryFn: () => getEthBalance(address),
	});

	if (query.status === "loading") {
		return (
			<div className="flex flex-col md:flex-row justify-between px-4 py-4 border-t">
				<div className="flex flex-row gap-4 items-center min-w-72">
					<Skeleton className="h-10 w-10 rounded-full" />
					<Skeleton className="h-4 w-[200px]" />
				</div>
				<div className="flex flex-row gap-4 items-center w-20">
					<Skeleton className="h-4 w-[50px]" />
				</div>
			</div>
		);
	}

	const eth = ethers.formatEther(query?.data || 0);

	return (
		<div className="flex flex-col md:flex-row justify-between gap-4 p-4 border-t">
			<div className="flex flex-row gap-4 items-center">
				<div className="bg-white rounded-full w-8 h-8 overflow-clip p-1 flex items-center place-content-center">
					<img
						src="/logos/sepolia.png"
						alt="Sepolia"
						className="w-auto h-6"
					/>
				</div>
				<div className="flex flex-col">
					<span className="text-muted-foreground text-xs">
						Sepolia Ether
					</span>
					<span>{eth} ETH</span>
				</div>
			</div>
			<div>
				<Button
					size="sm"
					variant="default"
					className="gap-2 text-sm w-28"
					disabled={eth === "0.0"}
				>
					<Link
						className="flex flex-row gap-4 items-center"
						to={`/new-transaction?key=${keyId}`}
					>
						<MoveUpRight className="h-4 w-4" />
						Send
					</Link>
				</Button>
			</div>
		</div>
	);
}
