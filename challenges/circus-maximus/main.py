from hashlib import sha256
import json
import time
import copy

class Block:
    def __init__(self, index, transactions, previous_hash):
        self.index = index
        self.transactions = transactions
        self.previous_hash = previous_hash
        self.nonce = 0

    def compute_hash(self):
        """
        A function that return the hash of the block contents.
        """
        block_string = json.dumps(self.__dict__, sort_keys=True)
        return sha256(block_string.encode()).hexdigest()

    def display_block(self):
        print("Block :" + str(self.index))
        print(self.transactions)


class Blockchain:
    # difficulty of our PoW algorithm
    difficulty = 2

    def __init__(self):
        self.unconfirmed_transactions = []
        self.chain = []

    def create_genesis_block(self):
        """
        A function to generate genesis block and appends it to
        the chain. The block has index 0, previous_hash as 0, and
        a valid hash.
        """
        genesis_block = Block(0, [], "0")
        genesis_block.hash = genesis_block.compute_hash()
        self.chain.append(genesis_block)

    @property
    def last_block(self):
        return self.chain[-1]

    def add_block(self, block, proof):
        """
        A function that adds the block to the chain after verification.
        Verification includes:
        * Checking if the proof is valid.
        * The previous_hash referred in the block and the hash of latest block
          in the chain match.
        """
        previous_hash = self.last_block.hash

        if previous_hash != block.previous_hash:
            print("Previous Hash Problem")
            return False

        if not Blockchain.is_valid_proof(block, proof):
            print("Unacceptable hash")
            return False

        block.hash = proof
        self.chain.append(block)
        return True

    @staticmethod
    def proof_of_work(block):
        """
        Function that tries different values of nonce to get a hash
        that satisfies our difficulty criteria.
        """
        block.nonce = 0

        computed_hash = block.compute_hash()
        while not computed_hash.startswith('0' * Blockchain.difficulty):
            block.nonce += 1
            computed_hash = block.compute_hash()

        return computed_hash

    def add_new_transaction(self, transaction_type, node_type, count_of_product):
        transaction = {}
        transaction["type"] = transaction_type
        transaction["node_type"] = node_type
        transaction["count_of_product"] = count_of_product

        self.unconfirmed_transactions.append(transaction)

    @classmethod
    def is_valid_proof(cls, block, block_hash):
        """
        Check if block_hash is valid hash of block and satisfies
        the difficulty criteria.
        """
        return (block_hash.startswith('0' * Blockchain.difficulty) and
                block_hash == block.compute_hash())

    @classmethod
    def check_chain_validity(cls, chain):
        result = True
        previous_hash = "0"

        for block in chain:
            block_hash = block.hash
            # remove the hash field to recompute the hash again
            # using `compute_hash` method.
            delattr(block, "hash")

            if not cls.is_valid_proof(block, block_hash) or \
                    previous_hash != block.previous_hash:
                result = False
                break

            block.hash, previous_hash = block_hash, block_hash

        return result
  
    def mine(self):
        """
        This function serves as an interface to add the pending
        transactions to the blockchain by adding them to the block
        and figuring out Proof Of Work.
        """
        if not self.unconfirmed_transactions:
            return False

        last_block = self.last_block

        new_block = Block(index=last_block.index + 1,
                          transactions=self.unconfirmed_transactions,
                          previous_hash=last_block.hash)

        proof = self.proof_of_work(new_block)
        self.add_block(new_block, proof)

        self.unconfirmed_transactions = []

        return True


    def consensus(self, peers):
        """
        Our naive consensus algorithm. If a longer valid chain is
        found, our chai:
            self = longest_chainn is replaced with it.
        """

        longest_chain = None
        current_len = len(self.chain)

        for node in peers:
            length = len(node.blockchain.chain)
            chain = node.blockchain
            if length > current_len and Blockchain.check_chain_validity(chain):
                current_len = length
                longest_chain = chain

        if longest_chain:
            self = longest_chain
            return True

        return False

    def announce_new_block(self, block, peers):
        """
        A function to announce to the network once a block has been mined.
        Other blocks can simply verify the proof of work and add it to their
        respective chains.
        """

        for peer in peers:
            block_new = Block(block.index, block.transactions, block.previous_hash)
            block_new.nonce = block.nonce
            proof = block.hash
            added = peer.blockchain.add_block(block_new, proof)

            if not added:
                return "The block was discarded by the node"

            return "Block added to the chain"


    def mine_and_announce(self, peers):
        result = self.mine()
        block = self.last_block

        if not result:
            return "No transactions to mine"
        else:
            for peer in peers:
                block_new = Block(block.index, block.transactions, block.previous_hash)
                block_new.nonce = block.nonce
                proof = block.hash
                peer.blockchain.add_block(block_new, proof)

            return "Block #{} is mined.".format(self.last_block.index)


    def display_chain(self):
        chain_data = []
        for block in self.chain:
            chain_data.append(block.__dict__)
            block.display_block()


class Node:
    def __init__(self, xcoord, ycoord):
        self.number_of_goods = 0
        self.capacity = 50
        self.blockchain = Blockchain()
        self.peers = []
        self.blockchain.create_genesis_block()
        self.xcoord = xcoord
        self.ycoord = ycoord

    def produce(self, count_of_goods):
        self.number_of_goods += count_of_goods
    
    def transaction(self, type_of_transaction, type_of_person, count_of_goods):
        self.blockchain.add_new_transaction(type_of_transaction, type_of_person, count_of_goods)
        if (type_of_transaction == "send"):
            self.number_of_goods -= count_of_goods
        if (type_of_transaction == "receive"):
            self.number_of_goods += count_of_goods

    def mine(self):
        self.blockchain.mine_and_announce(self.peers)
        
    def register_node(self, other_node):
        other_node.blockchain = copy.deepcopy(self.blockchain)
        for peer in self.peers:
            other_node.peers.append(peer)
        other_node.peers.append(self)
        self.peers.append(other_node)
        
    def display_chain(self):
        self.blockchain.display_chain()



def init_network():
    # Creating SCN nodes and specifying their location
    
    manufacturer1 = Node(30, 25)
    manufacturer2 = Node(10, 40)

    wholesaler1 = Node(10, 50)
    wholesaler2 = Node(50, 65)
    wholesaler3 = Node(27, 55)

    retailer1 = Node(50, 76)
    retailer2 = Node(25, 89)

    # Registering Peers
    manufacturer1.register_node(manufacturer2)
    manufacturer1.register_node(wholesaler1)
    manufacturer1.register_node(wholesaler2)
    manufacturer1.register_node(wholesaler3)
    manufacturer1.register_node(retailer1)
    manufacturer1.register_node(retailer2)

    manufacturer1.produce(20)
    manufacturer2.produce(30)
    manufacturer1.transaction("send", "wholesaler2", 4)
    manufacturer1.mine()
    manufacturer2.transaction("send", "wholesaler3", 5)
    manufacturer2.mine()
    wholesaler1.transaction("send", "retailer2", 3)
    manufacturer1.mine()

    print(manufacturer2.blockchain.last_block.hash)
    print(manufacturer2.blockchain.last_block.compute_hash())

init_network()