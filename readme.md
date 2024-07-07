# Consensus-Breaking Change

## What does it mean by breaking consensus?

Breaking consensus refers to making changes to the blockchain's protocol that are not backwards-compatible. When consensus is broken, nodes running the old version of the software will not agree with nodes running the new version, potentially leading to a chain split.

## Why this change breaks consensus

In this project, we modified how post IDs are generated. Previously, post IDs were assigned sequentially (count + 1). In the new version, we changed it to (count + 2).

This breaks consensus because:

1. Nodes running the old software will generate different post IDs than nodes running the new software for the same transactions.
2. This discrepancy in ID generation will lead to different state outcomes for the same set of transactions.
3. As a result, nodes will disagree on the state of the blockchain, breaking consensus.

To implement this change safely, a coordinated upgrade of all nodes would be required, along with a migration of existing data to the new format.
