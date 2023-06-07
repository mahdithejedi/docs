[awesome-postgres](https://github.com/dhamaniasad/awesome-postgres)

#Points
The file, postgresql.conf can be edited to increase the frequency of checkpoints, by reducing the checkpoint_segments and the checkpoint_timeout parameters. By reducing the integer value of these parameters, you will be flushing dirty pages more frequently, which is not desirable because checkpoints are costly. But on the other hand, after a crash, the recovery will be quicker.
<small>
    [Source](https://hevodata.com/learn/working-with-postgres-wal/#cp)
</small>
